/*
@Time : 2020/1/3 13:58
@Software: GoLand
@File : init
@Author : https://github.com/hsyan2008/gom
*/
package Sql2struct

import (
	"encoding/json"
	"github.com/buger/jsonparser"
	_app "gormat/app"
	"log"
	"xorm.io/core"
)

var Tables []*core.Table

type SQL2Struct struct {
	TargetDir     string   `json:"target_dir"`
	AutoSave      bool     `json:"auto_save"`
	Tags          []string `json:"tags"`
	JSONOmitempty bool     `json:"json_omitempty"`
	Driver        string   `json:"driver"`
	ExcludeTables []string `json:"exclude_tables"`
	TryComplete   bool     `json:"try_complete"`
	Tinyint2Bool  bool     `json:"tinyint2bool"`
	SourceMap     struct {
		User     string `json:"user"`
		Password string `json:"password"`
		Host     string `json:"host"`
		Port     string `json:"port"`
		Db       string `json:"db"`
	} `json:"sourceMap"`
	Reflect string `json:"reflect"`
	Special string `json:"special"`
}

func Configs() (config *SQL2Struct) {
	data, _, _, _ := jsonparser.Get(_app.Config, "sql2struct")
	if err := json.Unmarshal([]byte(data), &config); err != nil {
		log.Println(err.Error())
	}
	return
}
