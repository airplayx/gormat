package Sql2struct

import (
	"strings"

	"xorm.io/core"
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

func NewModel(table *core.Table, maps map[string]string) (m model) {
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

func NewModelField(table *core.Table, column *core.Column, maps map[string]string) (f modelField) {
	f = modelField{
		FieldName:  core.LintGonicMapper.Table2Obj(column.Name),
		ColumnName: column.Name,
		Type:       maps[strings.ToLower(column.SQLType.Name)],
		Imports:    getGoImports(column),
	}
	if strings.HasPrefix(f.ColumnName, "is_") && column.SQLType.Name == core.TinyInt {
		f.Type = "bool"
	}
	for key, val := range JSONMethod(Configs().Special) {
		if f.ColumnName == key {
			f.Type = val
		}
	}
	var tags []string
	for _, v := range Configs().Tags {
		switch v {
		case "json":
			tags = append(tags, GetJsonTag(column, Configs().JSONOmitempty))
		case "gorm":
			tags = append(tags, GetGormTag(table, column))
		}
	}
	if len(tags) > 0 {
		f.Tag = "`" + strings.Join(tags, " ") + "`"
	}
	return
}
