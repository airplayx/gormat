package Sql2struct

import (
	"encoding/json"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"xorm.io/core"
)

type GenTool struct {
	targetDir   string
	packageName string
	tables      []*core.Table
	models      map[string]model
}

func NewGenTool() *GenTool {
	dir := Configs().TargetDir
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		log.Println(err.Error())
	}
	return &GenTool{
		targetDir:   dir,
		packageName: filepath.Base(dir),
		models:      make(map[string]model),
	}
}

func (genTool *GenTool) getDBMetas(ts []string) (err error) {
	genTool.tables, err = DBMetas(ts, Configs().ExcludeTables, Configs().TryComplete)
	if err != nil {
		return
	}

	return nil
}

func (genTool *GenTool) genModels(maps string) {
	for _, table := range genTool.tables {
		model := NewModel(table, maps)
		genTool.models[model.TableName] = model
	}
	return
}

func (genTool *GenTool) genFile() (by []byte, err error) {
	for _, model := range genTool.models {
		//package
		str := fmt.Sprintln("package", genTool.packageName)

		//import
		if len(model.Imports) > 0 {
			str += fmt.Sprintln("import (")
			for _, i := range model.Imports {
				str += fmt.Sprintln(`"` + i + `"`)
			}
			str += fmt.Sprintln(")")
		}
		//struct
		str += fmt.Sprintln("type", model.StructName, "struct {")
		for _, v := range model.Fields {
			str += fmt.Sprintln(v.FieldName, v.Type, v.Tag, v.Comment)
		}
		str += fmt.Sprintln("}")

		//func
		str += fmt.Sprintln("func (*", model.StructName, ") TableName() string {")
		str += fmt.Sprintln(fmt.Sprintf("return `%s` //"+model.Comment, model.TableName))
		str += fmt.Sprintln("}")
		//format
		by, err = format.Source([]byte(str))
		if err != nil {
			return
		}
		if Configs().AutoSave {
			file := filepath.Join(genTool.targetDir, fmt.Sprintf("%s.go", model.TableName))
			if err = ioutil.WriteFile(file, by, 0644); err != nil {
				return
			}
			log.Println("gen into file:", file)
		}
	}
	return
}

func (genTool *GenTool) Gen(ts []string, dbConf *SourceMap) (result []byte, err error) {
	if err = InitDb(dbConf); err != nil {
		return
	}
	if err = genTool.getDBMetas(ts); err != nil {
		return
	}
	reflect, _ := json.Marshal(Configs().Reflect)
	var config string
	if err = json.Unmarshal(reflect, &config); err != nil {
		return
	}
	genTool.genModels(config)
	return genTool.genFile()
}
