package sql2struct

import (
	"encoding/json"
	"strings"

	"github.com/xormplus/core"
)

//Model ...
type Model struct {
	StructName string
	TableName  string
	Imports    map[string]string
	Fields     []ModelField
	Comment    string
}

//ModelField ...
type ModelField struct {
	FieldName  string
	ColumnName string
	Type       string
	Imports    map[string]string
	Tag        string
	Comment    string
}

//NewModelField ...
func NewModelField(table *core.Table, column *core.Column, maps string) (f *ModelField) {
	var reflect, special map[string]string
	_ = json.Unmarshal([]byte(maps), &reflect)
	f = &ModelField{
		FieldName:  core.LintGonicMapper.Table2Obj(column.Name),
		ColumnName: column.Name,
		Type:       reflect[strings.ToLower(column.SQLType.Name)],
		Imports:    getGoImports(column),
	}
	f.Type = getTypeAndImports(column)

	s, _ := json.Marshal(Configs().Special)
	var config string
	if err := json.Unmarshal(s, &config); err != nil {
		return
	}
	_ = json.Unmarshal([]byte(config), &special)
	for key, val := range special {
		if f.ColumnName == key {
			f.Type = val
		}
	}
	var tags []string
	for _, v := range Configs().Tags {
		switch v {
		case "json":
			tags = append(tags, GetJSONTag(column, Configs().JSONOmitempty))
		case "xorm":
			tags = append(tags, GetXormTag(table, column))
		case "gorm":
			tags = append(tags, GetGormTag(table, column))
		}
	}
	if len(tags) > 0 {
		f.Tag = "`" + strings.Join(tags, " ") + "`"
	}
	return
}
