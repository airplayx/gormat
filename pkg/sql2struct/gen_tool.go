package sql2struct

import (
	"encoding/json"
	"fmt"
	"github.com/xormplus/core"
	"github.com/xormplus/xorm/schemas"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

//GenTool ...
type GenTool struct {
	targetDir   string
	packageName string
	tables      []*schemas.Table
	models      map[string]Model
}

//NewGenTool ...
func NewGenTool() *GenTool {
	dir := Configs().TargetDir
	if Configs().AutoSave {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			log.Println(err.Error())
		}
	}
	return &GenTool{
		targetDir:   dir,
		packageName: filepath.Base(dir),
		models:      make(map[string]Model),
	}
}

//getDBMetas ...
func (genTool *GenTool) getDBMetas() (err error) {
	genTool.tables, err = Engine.DBMetas()
	if err != nil {
		return
	}
	return nil
}

func (genTool *GenTool) genFile(table *schemas.Table) (by []byte, err error) {
	if v, ok := genTool.models[table.Name]; ok {
		//package
		str := fmt.Sprintln("package", genTool.packageName)
		//import
		if len(v.Imports) > 0 {
			str += fmt.Sprintln("import (")
			for _, i := range v.Imports {
				str += fmt.Sprintln(`"` + i + `"`)
			}
			str += fmt.Sprintln(")")
		}
		//struct
		str += fmt.Sprintln("type", v.StructName, "struct {")
		for _, f := range v.Fields {
			str += fmt.Sprintln(f.FieldName, f.Type, f.Tag, f.Comment)
		}
		str += fmt.Sprintln("}")

		//func
		str += fmt.Sprintln("func (f *", v.StructName, ") TableName() string {")
		str += fmt.Sprintln(fmt.Sprintf("return `%s` //"+v.Comment, v.TableName))
		str += fmt.Sprintln("}")

		//setFunc
		for _, f := range v.Fields {
			str += fmt.Sprintln(fmt.Sprintf("\nfunc (f *%s) Set%s(n %s) {", v.StructName, f.FieldName, f.Type))
			str += fmt.Sprintln(fmt.Sprintf("f.%s = n", f.FieldName))
			str += fmt.Sprintln("}")
		}
		//format
		by, err = format.Source([]byte(str))
		if err != nil {
			return
		}
		if Configs().AutoSave {
			file := filepath.Join(genTool.targetDir, fmt.Sprintf("%s.go", v.TableName))
			if err = ioutil.WriteFile(file, by, 0644); err != nil {
				return
			}
			log.Println("gen into file:", file)
			if Configs().SaveFactoryFunc {
				if err := saveFacFile(genTool.targetDir, genTool.packageName, v); err != nil {
					return by, err
				}
			}
		}
	}
	return
}

func saveFacFile(targetDir, packageName string, v Model) error {
	facStr := fmt.Sprintln("package", packageName)
	if len(v.Imports) > 0 {
		facStr += fmt.Sprintln("import (")
		for _, i := range v.Imports {
			facStr += fmt.Sprintln(`"` + i + `"`)
		}
		facStr += fmt.Sprintln(")")
	}

	//factory
	facStr += fmt.Sprintln(fmt.Sprintf("type %sFactory interface {", v.StructName))
	for _, f := range v.Fields {
		facStr += fmt.Sprintln(fmt.Sprintf("Set%s(%s)", f.FieldName, f.Type))
	}
	facStr += fmt.Sprintln(fmt.Sprintf("Get%s() %s", v.StructName, v.StructName))
	for _, f := range v.Fields {
		facStr += fmt.Sprintln(fmt.Sprintf("Get%s() %s", f.FieldName, f.Type))
	}
	facStr += fmt.Sprintln("}")

	//extension
	facStr += fmt.Sprintln(fmt.Sprintf("\ntype %sExtension struct {", v.StructName))
	facStr += fmt.Sprintln(fmt.Sprintf("*%s", v.StructName))
	facStr += fmt.Sprintln("}")
	facStr += fmt.Sprintln(fmt.Sprintf("\nfunc (f %sExtension) Get%s() %s {", v.StructName, v.StructName, v.StructName))
	facStr += fmt.Sprintln(fmt.Sprintf("return *f.%s", v.StructName))
	facStr += fmt.Sprintln("}")
	for _, f := range v.Fields {
		facStr += fmt.Sprintln(fmt.Sprintf("\nfunc (f %sExtension) Get%s() %s {", v.StructName, f.FieldName, f.Type))
		facStr += fmt.Sprintln(fmt.Sprintf("return f.%s", f.FieldName))
		facStr += fmt.Sprintln("}")
	}
	//format
	funcData, _ := format.Source([]byte(facStr))
	facFile := filepath.Join(targetDir, fmt.Sprintf("%s_factory.go", v.TableName))
	if err := ioutil.WriteFile(facFile, funcData, 0644); err != nil {
		return err
	}
	log.Println("gen into file:", facFile)
	return nil
}

//Gen ...
func (genTool *GenTool) Gen(table *schemas.Table, dbConf *SourceMap) (result []byte, err error) {
	//if err = Init(dbConf); err != nil {
	//	return
	//}

	//if err = genTool.getDBMetas(); err != nil {
	//	return
	//}

	reflect, _ := json.Marshal(Configs().Reflect)
	var config string
	if err = json.Unmarshal(reflect, &config); err != nil {
		return
	}
	//for _, table := range genTool.tables {
	m := &Model{
		StructName: core.LintGonicMapper.Table2Obj(table.Name),
		TableName:  table.Name,
		Imports:    map[string]string{},
		Comment:    table.Comment,
	}
	for _, column := range table.Columns() {
		f := NewModelField(table, column, config)
		for k, v := range f.Imports {
			m.Imports[k] = v
		}
		m.Fields = append(m.Fields, *f)
	}
	genTool.models[m.TableName] = *m
	//}
	return genTool.genFile(table)
}
