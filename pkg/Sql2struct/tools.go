package Sql2struct

import (
	"github.com/xormplus/core"
	"strings"
)

func getTypeAndImports(column *core.Column) (t string) {
	t = sqlType2TypeString(column.SQLType)
	if Configs().Tinyint2bool && strings.HasPrefix(column.Name, "is_") &&
		column.SQLType.Name == "TINYINT" && column.SQLType.DefaultLength == 1 {
		t = "bool"
		return
	}
	return
}

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

func InStringSlice(f string, a []string) bool {
	for _, s := range a {
		if f == s {
			return true
		}
	}
	return false
}
