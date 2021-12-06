package sql2struct

import (
	"fmt"
	"github.com/xormplus/xorm/schemas"
	"sort"
	"strings"

	"github.com/xormplus/core"
)

//GetGormTag ...
func GetGormTag(table *schemas.Table, col *schemas.Column) string {
	isNameID := col.Name == table.AutoIncrement
	isIDPk := isNameID && sqlType2TypeString(col.SQLType) == "int64"

	var res []string
	res = append(res, "column:"+col.Name)

	if !col.Nullable {
		if !isIDPk {
			res = append(res, "not null")
		}
	}
	if col.IsPrimaryKey {
		res = append(res, "not null;primary_key")
	}
	if col.Default != "" {
		def := strings.Trim(col.Default, "'")
		if strings.EqualFold(def, "NULL") {
			def = "null"
		} else if def == "current_timestamp()" || def == "CURRENT_TIMESTAMP" {
			def = "current_timestamp()"
		} else {
			def = "'" + def + "'"
		}
		res = append(res, "default:"+def+"")
	}
	if col.IsAutoIncrement {
		res = append(res, "AUTO_INCREMENT")
	}

	names := make([]string, 0, len(col.Indexes))
	for name := range col.Indexes {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		index := table.Indexes[name]
		var s string
		if index.Type == core.UniqueType {
			if len(index.Cols) > 1 {
				s += "unique_index:" + index.Name
			} else {
				s = "unique"
			}
		} else if index.Type == core.IndexType {
			s = "index"
			if len(index.Cols) > 1 {
				s += ":" + index.Name
			}
		}
		res = append(res, s)
		var result []string
		tempMap := map[string]byte{}
		for _, e := range res {
			l := len(tempMap)
			tempMap[e] = 0
			if len(tempMap) != l {
				result = append(result, e)
			}
		}
		res = result
	}

	s := "type:" + strings.ToLower(col.SQLType.Name)
	if col.Length != 0 {
		if col.Length2 != 0 {
			s += fmt.Sprintf("(%v,%v)", col.Length, col.Length2)
		} else {
			s += fmt.Sprintf("(%v)", col.Length)
		}
	} else if len(col.EnumOptions) > 0 { //enum
		s += "("
		opts := ""

		enumOptions := make([]string, 0, len(col.EnumOptions))
		for enumOption := range col.EnumOptions {
			enumOptions = append(enumOptions, enumOption)
		}
		sort.Strings(enumOptions)

		for _, v := range enumOptions {
			opts += fmt.Sprintf(",'%v'", v)
		}
		s += strings.TrimLeft(opts, ",")
		s += ")"
	} else if len(col.SetOptions) > 0 { //enum
		s += "("
		opts := ""

		setOptions := make([]string, 0, len(col.SetOptions))
		for setOption := range col.SetOptions {
			setOptions = append(setOptions, setOption)
		}
		sort.Strings(setOptions)

		for _, v := range setOptions {
			opts += fmt.Sprintf(",'%v'", v)
		}
		s += strings.TrimLeft(opts, ",")
		s += ")"
	}
	if col.Comment != "" {
		res = append(res, "comment:'"+col.Comment+"'")
	}

	if !col.IsPrimaryKey {
		res = append(res, s)
	}
	if len(res) > 0 {
		return "gorm:\"" + strings.Join(res, ";") + "\""
	}

	return ""
}
