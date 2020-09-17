/*
@Time : 2020/1/15 15:07
@Software: GoLand
@File : init
@Author : https://github.com/yujiahaol68/sql2struct
*/

package quickly

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"gormat/pkg/sql2struct"
	"io"
	"regexp"
	"sort"
	"strings"
)

var (
	tableStmt = regexp.MustCompile(`(?is)create table (\w+) \(\s.+?;`)
	colDecl   = regexp.MustCompile(`\w.+`)

	upperDict map[string]struct{}
	metaTag   = sql2struct.Configs().Tags
)

type column struct {
	Name       []byte
	Constraint []byte
}

//QuickOut ...
type QuickOut struct {
	Name   string
	Fields map[int]field
}

type field struct {
	Name      string
	FieldType string
	MetaInfo  map[string][]string
}

func init() {
	upperDict = map[string]struct{}{}
	upperDict["id"] = struct{}{}
	upperDict["url"] = struct{}{}
}

func (f field) String() string {
	meta := make([]string, 0, len(f.MetaInfo))
	for _, tag := range metaTag {
		if v, ok := f.MetaInfo[tag]; ok {
			meta = append(meta, fmt.Sprintf(`%s:"%s"`, tag, strings.Join(v, "")))
		}
	}

	var fieldName string
	if _, ok := upperDict[f.Name]; ok {
		fieldName = strings.ToUpper(f.Name)
	} else {
		fieldName = strings.Title(f.Name)
	}

	return fmt.Sprintf("%s %s `%s`", fieldName, f.FieldType, strings.Join(meta, " "))
}

//GenType ...
func (t QuickOut) GenType() ([]byte, error) {
	str := fmt.Sprintln("type", t.Name, "struct {")
	sortedMap(t.Fields, func(k int, v field) {
		str += fmt.Sprintln(v)
	})
	str += fmt.Sprintln("}")
	return format.Source([]byte(str))
}

func sortedMap(m map[int]field, f func(k int, v field)) {
	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		f(k, m[k])
	}
}

func newField(name, constraint []byte) (f field) {
	f = field{
		Name:     toCamel(bytes.TrimLeft(name, "_")),
		MetaInfo: map[string][]string{},
	}
	col := strings.ToUpper(string(constraint))
	f.FieldType, _ = scanType(name, strings.ToLower(strings.Split(col, " ")[0]))
	for _, v := range sql2struct.Configs().Tags {
		switch v {
		case "gorm":
			tag := []string{"column:" + string(name) + ";"}
			if strings.Contains(col, "NOT NULL") {
				tag = append(tag, "not null;")
			}
			if strings.Contains(col, "DEFAULT") {
				tag = append(tag, "default:"+strings.Split(strings.Split(col, "DEFAULT")[1], " ")[1]+";")
			}
			if strings.Contains(col, "COMMENT") {
				tag = append(tag, "comment:"+strings.Split(col, "COMMENT")[1]+";")
			}
			if strings.Contains(col, "AUTO_INCREMENT") {
				tag = append(tag, "AUTO_INCREMENT;")
			}
			tag = append(tag, "type:"+strings.ToLower(strings.Split(col, " ")[0]))
			f.MetaInfo[v] = tag
		case "json":
			if sql2struct.Configs().JSONOmitempty {
				f.MetaInfo[v] = []string{string(name), ",omitempty"}
			} else {
				f.MetaInfo[v] = []string{string(name)}
			}
		default:
			f.MetaInfo[v] = []string{string(name)}
		}
	}
	return
}

func toCamel(s []byte) string {
	if !bytes.Contains(s, []byte("_")) {
		return string(s)
	}
	return hunToCamel(string(s))
}

func scanType(name []byte, tag string) (s string, haveType bool) {
	var reflect, special map[string]string
	_ = json.Unmarshal([]byte(sql2struct.Configs().Reflect), &reflect)
	s = reflect[strings.Split(tag, "(")[0]]
	_ = json.Unmarshal([]byte(sql2struct.Configs().Special), &special)
	for key, val := range special {
		if string(name) == key {
			s = val
		}
	}
	haveType = s != ""
	return
}

//MatchStmt ...
func MatchStmt(r io.Reader) (byte [][][]byte, err error) {
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(r)
	byte = tableStmt.FindAllSubmatch(buf.Bytes(), -1)
	return
}

//HandleStmtBlock ...
func HandleStmtBlock(s [][]byte) (t QuickOut) {
	block := s[0]
	leftTrimIdx := 0
	rightTrimIdx := len(block) - 1
	for ; leftTrimIdx < len(block) && block[leftTrimIdx] != '('; leftTrimIdx++ {
	}
	for ; rightTrimIdx >= 0 && block[rightTrimIdx] != ')'; rightTrimIdx-- {
	}
	block = block[leftTrimIdx+1 : rightTrimIdx]
	cols := extractRawCols(block)
	t.Name = strings.Title(toCamel(s[1]))
	t.Fields = make(map[int]field)
	for k, v := range cols {
		f := newField(v.Name, v.Constraint)
		if f.FieldType == "" {
			continue
		}
		t.Fields[k] = f
	}
	return t
}

func extractRawCols(s []byte) []column {
	cols := colDecl.FindAll(s, -1)
	allColumns := make([]column, len(cols))

	for k, v := range cols {
		v = bytes.TrimRight(v, ", ")
		c := bytes.SplitN(v, []byte{' '}, 2)

		allColumns[k].Name = c[0]
		allColumns[k].Constraint = bytes.ToLower(c[1])
	}
	return allColumns
}

func hunToCamel(str string) string {
	s := strings.Split(str, "_")
	var ns string
	for i := range s {
		if i == 0 {
			ns += s[i]
			continue
		}
		if _, ok := upperDict[s[i]]; ok {
			ns += strings.ToUpper(s[i])
			continue
		}
		ns += strings.Title(s[i])
	}
	return ns
}
