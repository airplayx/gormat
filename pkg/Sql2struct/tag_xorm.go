/*
@Time : 2020/1/7 11:04
@Software: GoLand
@File : xorm_tag
@Author : Bingo <airplayx@gmail.com>
*/
package Sql2struct

import (
	"sort"
	"strings"

	"github.com/xormplus/core"
)

var created = []string{"created_at"}
var updated = []string{"updated_at"}
var deleted = []string{"deleted_at"}

func GetXormTag(table *core.Table, col *core.Column) string {
	isNameId := col.Name == table.AutoIncrement
	isIdPk := isNameId && sqlType2TypeString(col.SQLType) == "int64"

	var res []string
	if !col.Nullable {
		if !isIdPk {
			res = append(res, "not null")
		}
	}

	if col.IsPrimaryKey {
		res = append(res, "pk")
	}

	if len(col.Default) >= 4 && strings.HasPrefix(col.Default, "''") && strings.HasSuffix(col.Default, "''") {
		col.Default = col.Default[1 : len(col.Default)-1]
	}
	if col.Default != "" {
		res = append(res, "default "+col.Default)
	}

	if col.IsAutoIncrement {
		res = append(res, "autoincr")
	}

	if col.SQLType.IsTime() && InStringSlice(col.Name, created) {
		res = append(res, "created")
	}

	if col.SQLType.IsTime() && InStringSlice(col.Name, updated) {
		res = append(res, "updated")
	}

	if col.SQLType.IsTime() && InStringSlice(col.Name, deleted) {
		res = append(res, "deleted")
	}

	names := make([]string, 0, len(col.Indexes))
	for name := range col.Indexes {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		index := table.Indexes[name]
		var uistr string
		if index.Type == core.UniqueType {
			uistr = "unique"
		} else if index.Type == core.IndexType {
			uistr = "index"
		}
		if len(index.Cols) > 1 {
			uistr += "(" + index.Name + ")"
		}
		res = append(res, uistr)
	}

	res = append(res, DB().SQLType(col))

	if len(res) > 0 {
		return "xorm:\"" + strings.Join(res, " ") + "\""
	}

	return ""
}
