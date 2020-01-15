/*
@Time : 2019/12/25 16:05
@Software: GoLand
@File : config
@Author : Bingo <airplayx@gmail.com>
*/
package config

import (
	"io/ioutil"
	"os"
)

var (
	Setting = []byte(`{"const":{"font":"./miniHei.ttf","theme":"light","scale":1},"sql2struct":{"target_dir":"./models","auto_save":false,"tags":["gorm","json"],"json_omitempty":true,"driver":"mysql","exclude_tables":[],"try_complete":true,"tinyint2bool":true,"sourceMap":[],"reflect":"{\"tinyint\":\"int8\",\"smallint\":\"int16\",\"int\":\"int32\",\"bigint\":\"int64\",\"float\":\"float64\",\"double\":\"float64\",\"decimal\":\"float64\",\"char\":\"string\",\"varchar\":\"string\",\"text\":\"string\",\"mediumtext\":\"string\",\"longtext\":\"string\",\"time\":\"time.Time\",\"date\":\"time.Time\",\"datetime\":\"time.Time\",\"timestamp\":\"time.Time\",\"enum\":\"string\",\"set\":\"string\",\"blob\":\"string\"}","special":"{\"id\":\"uint\"}"}}`)
	File    = "./config.json"
)

func init() {
	_, err := os.Stat(File)
	if err == nil {
		Setting, _ = ioutil.ReadFile(File)
	}
}
