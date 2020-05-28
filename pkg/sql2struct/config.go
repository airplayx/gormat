/*Package sql2struct ...
@Time : 2020/1/3 13:58
@Software: GoLand
@File : init
@Author : https://github.com/hsyan2008/gom
*/
package sql2struct

import (
	"encoding/json"
	"github.com/buger/jsonparser"
	"gormat/configs"
	"log"
)

//SQL2Struct ...
type SQL2Struct struct {
	AutoSave      bool        `json:"auto_save"`
	JSONOmitempty bool        `json:"json_omitempty"`
	Reflect       string      `json:"reflect"`
	SourceMap     []SourceMap `json:"sourceMap"`
	Special       string      `json:"special"`
	Tags          []string    `json:"tags"`
	TargetDir     string      `json:"target_dir"`
	Tinyint2bool  bool        `json:"tinyint2bool"`
}

//SourceMap ...
type SourceMap struct {
	Db       []string `json:"db"`
	Driver   string   `json:"driver"`
	Host     string   `json:"host"`
	Password string   `json:"password"`
	Port     string   `json:"port"`
	User     string   `json:"user"`
}

//Configs ...
func Configs() (s2s *SQL2Struct) {
	data, _, _, _ := jsonparser.Get(configs.JSON, "sql2struct")
	if err := json.Unmarshal([]byte(data), &s2s); err != nil {
		log.Println(err.Error())
	}
	return
}
