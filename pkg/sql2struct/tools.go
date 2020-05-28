package sql2struct

import (
	"github.com/xormplus/xorm/schemas"
	"strings"
)

func getTypeAndImports(column *schemas.Column) (t string) {
	t = sqlType2TypeString(column.SQLType)
	if Configs().Tinyint2bool && strings.HasPrefix(column.Name, "is_") &&
		column.SQLType.Name == "TINYINT" && column.SQLType.DefaultLength == 1 {
		t = "bool"
		return
	}
	return
}

//RmDuplicateElement ...
func RmDuplicateElement(keywords []string) []string {
	result := make([]string, 0, len(keywords))
	temp := map[string]struct{}{}
	for _, item := range keywords {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

func inStringSlice(f string, a []string) bool {
	for _, s := range a {
		if f == s {
			return true
		}
	}
	return false
}

func sqlType2TypeString(st schemas.SQLType) string {
	t := schemas.SQLType2Type(st)
	s := t.String()
	if s == "[]uint8" {
		return "[]byte"
	}
	return s
}

func getGoImports(column *schemas.Column) map[string]string {
	imports := make(map[string]string)

	if sqlType2TypeString(column.SQLType) == "time.Time" {
		imports["time"] = "time"
	}

	return imports
}
