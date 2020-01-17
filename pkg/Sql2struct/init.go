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
	"gormat/configs"
	"log"
)

type SQL2Struct struct {
	AutoSave      bool        `json:"auto_save"`
	ExcludeTables []string    `json:"exclude_tables"`
	JsonOmitempty bool        `json:"json_omitempty"`
	Reflect       string      `json:"reflect"`
	SourceMap     []SourceMap `json:"sourceMap"`
	Special       string      `json:"special"`
	Tags          []string    `json:"tags"`
	TargetDir     string      `json:"target_dir"`
	Tinyint2bool  bool        `json:"tinyint2bool"`
	TryComplete   bool        `json:"try_complete"`
}

type SourceMap struct {
	Db       []string `json:"db"`
	Driver   string   `json:"driver"`
	Host     string   `json:"host"`
	Password string   `json:"password"`
	Port     string   `json:"port"`
	User     string   `json:"user"`
}

func Configs() (s2s *SQL2Struct) {
	data, _, _, _ := jsonparser.Get(configs.Json, "sql2struct")
	if err := json.Unmarshal([]byte(data), &s2s); err != nil {
		log.Println(err.Error())
	}
	return
}
