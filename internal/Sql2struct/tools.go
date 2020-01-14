package Sql2struct

import (
	"strings"
	"xorm.io/core"
)

func InStringSlice(f string, a []string) bool {
	for _, s := range a {
		if f == s {
			return true
		}
	}
	return false
}

func getTypeAndImports(column *core.Column) (t string) {
	t = sqlType2TypeString(column.SQLType)
	if Configs().Tinyint2bool && strings.HasPrefix(column.Name, "is_") &&
		column.SQLType.Name == "TINYINT" && column.SQLType.DefaultLength == 1 {
		t = "bool"
		return
	}
	return
}
