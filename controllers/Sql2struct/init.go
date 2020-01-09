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
)

type SQL2Struct struct {
	AutoSave      bool     `json:"auto_save"`
	Driver        string   `json:"driver"`
	ExcludeTables []string `json:"exclude_tables"`
	JsonOmitempty bool     `json:"json_omitempty"`
	Reflect       string   `json:"reflect"`
	SourceMap     []struct {
		Db       []string `json:"db"`
		Host     string   `json:"host"`
		Password string   `json:"password"`
		Port     string   `json:"port"`
		User     string   `json:"user"`
	} `json:"sourceMap"`
	Special      string   `json:"special"`
	Tags         []string `json:"tags"`
	TargetDir    string   `json:"target_dir"`
	Tinyint2bool bool     `json:"tinyint2bool"`
	TryComplete  bool     `json:"try_complete"`
}

func Configs() (config *SQL2Struct) {
	data, _, _, _ := jsonparser.Get(_app.Config, "sql2struct")
	if err := json.Unmarshal([]byte(data), &config); err != nil {
		log.Println(err.Error())
	}
	return
}
