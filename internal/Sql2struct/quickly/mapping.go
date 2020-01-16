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
	"gormat/internal/Sql2struct"
	"io"
	"regexp"
	"strings"
)

var (
	tableStmt = regexp.MustCompile(`(?is)create table (\w+) \(\s.+?;`)
	comment   = regexp.MustCompile(`--.*`)
	colDecl   = regexp.MustCompile(`\w.+`)

	intType    = regexp.MustCompile(`(int|integer|smallint|bigint|serial)`)
	boolType   = regexp.MustCompile(`boolean`)
	floatType  = regexp.MustCompile(`(float|decimal|numeric|real)`)
	timeType   = regexp.MustCompile(`(date|time|timestamp)`)
	stringType = regexp.MustCompile(`(char|varchar|text)`)

	hungaryStyle = regexp.MustCompile(`_[a-z]+`)
	camelStyle   = regexp.MustCompile(`[a-z][A-Z]`)

	upperDict map[string]struct{}
	metaTag   = Sql2struct.Configs().Tags
)

type column struct {
	Name       []byte
	Constraint []byte
}

type langT struct {
	Name   string
	Fields map[string]field
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

func (t langT) GenType(w io.Writer) {
	fmt.Fprintf(w, "\ntype %s struct {\n", t.Name)
	for _, f := range t.Fields {
		fmt.Fprintf(w, "\t%s\n", f)
	}
	fmt.Fprintln(w, "}")
}

func newField(name, constraint []byte) (f field) {
	f = field{
		Name:     toCamel(bytes.TrimLeft(name, "_")),
		MetaInfo: map[string][]string{},
	}
	col := strings.ToUpper(string(constraint))
	f.FieldType, _ = scanType(name, strings.ToLower(strings.Split(col, " ")[0]))
	for _, v := range Sql2struct.Configs().Tags {
		if v == "gorm" {
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
			continue
		}
		if v == "json" {
			if Sql2struct.Configs().JsonOmitempty {
				f.MetaInfo[v] = []string{string(name), ",omitempty"}
			} else {
				f.MetaInfo[v] = []string{string(name)}
			}
			continue
		}
		f.MetaInfo[v] = []string{string(name)}
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
	_ = json.Unmarshal([]byte(Sql2struct.Configs().Reflect), &reflect)
	s = reflect[strings.Split(tag, "(")[0]]
	_ = json.Unmarshal([]byte(Sql2struct.Configs().Special), &special)
	for key, val := range special {
		if string(name) == key {
			s = val
		}
	}
	haveType = s != ""
	return
}

func (t *langT) AddJSONinfo(fieldName, info string) {
	if f, ok := t.Fields[fieldName]; ok {
		if _, ok := f.MetaInfo["json"]; ok {
			f.MetaInfo["json"] = append(f.MetaInfo["json"], info)
		} else {
			f.MetaInfo["json"] = []string{info}
		}
	}
}

func MatchStmt(r io.Reader) (byte [][][]byte, err error) {
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(r)

	stmt := buf.Bytes()
	byte = tableStmt.FindAllSubmatch(stmt, -1)
	return
}

func HandleStmtBlock(s [][]byte) langT {
	block := s[0]
	leftTrimIdx := 0
	rightTrimIdx := len(block) - 1
	for ; leftTrimIdx < len(block) && block[leftTrimIdx] != '('; leftTrimIdx++ {
	}
	for ; rightTrimIdx >= 0 && block[rightTrimIdx] != ')'; rightTrimIdx-- {
	}
	block = block[leftTrimIdx+1 : rightTrimIdx]
	block = delComment(block)

	//fmt.Printf("%s\n***\n", string(block))
	cols := extractRawCols(block)
	var t langT
	t.Name = strings.Title(toCamel(s[1]))
	t.Fields = make(map[string]field)
	for k := range cols {
		f := newField(cols[k].Name, cols[k].Constraint)
		if f.FieldType == "" {
			continue
		}
		t.Fields[f.Name] = f
	}
	return t
}

func delComment(s []byte) []byte {
	return comment.ReplaceAll(s, nil)
}

func extractRawCols(s []byte) []column {
	cols := colDecl.FindAll(s, -1)
	allColumns := make([]column, len(cols))

	for i := range cols {
		cols[i] = bytes.TrimRight(cols[i], ", ")
		c := bytes.SplitN(cols[i], []byte{' '}, 2)

		allColumns[i].Name = c[0]
		allColumns[i].Constraint = bytes.ToLower(c[1])
	}
	//fmt.Printf("%q\n", allColumns)

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

func camelTohun(str []byte) string {
	pos := camelStyle.FindAllSubmatchIndex(str, -1)
	if len(pos) == 0 {
		return string(str)
	}

	underPos := make([]int, len(pos))
	for i := range pos {
		underPos[i] = pos[i][0]
	}

	str = bytes.ToLower(str)
	nb := make([]byte, 0, len(str)+len(pos))
	underIdx := 0
	for i := range str {
		nb = append(nb, str[i])
		if underIdx < len(underPos) && i == underPos[underIdx] {
			nb = append(nb, '_')
			underIdx++
		}
	}
	return string(nb)
}
