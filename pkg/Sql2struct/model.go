package Sql2struct

import (
	"encoding/json"
	"strings"

	"github.com/xormplus/core"
)

type model struct {
	StructName string
	TableName  string
	Imports    map[string]string
	Fields     []modelField
	Comment    string
}

type modelField struct {
	FieldName  string
	ColumnName string
	Type       string
	Imports    map[string]string
	Tag        string
	Comment    string
}

func NewModel(table *core.Table, maps string) (m model) {
	m = model{
		StructName: core.LintGonicMapper.Table2Obj(table.Name),
		TableName:  table.Name,
		Imports:    map[string]string{},
		Comment:    table.Comment,
	}
	for _, column := range table.Columns() {
		f := NewModelField(table, column, maps)
		for k, v := range f.Imports {
			m.Imports[k] = v
		}
		m.Fields = append(m.Fields, f)
	}
	return
}

func NewModelField(table *core.Table, column *core.Column, maps string) (f modelField) {
	var reflect, special map[string]string
	_ = json.Unmarshal([]byte(maps), &reflect)
	f = modelField{
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
			tags = append(tags, GetJsonTag(column, Configs().JsonOmitempty))
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
