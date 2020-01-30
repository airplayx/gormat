/*
@Time : 2019/12/25 16:05
@Software: GoLand
@File : init
@Author : Bingo <airplayx@gmail.com>
*/
package configs

import (
	"io/ioutil"
	"os"
)

var (
	Json = []byte(`{
  "const": {
    "font": "./miniHei.ttf",
    "theme": "light",
    "scale": "1.0"
  },
  "sql2struct": {
    "target_dir": "./models",
    "auto_save": false,
    "tags": [
      "gorm",
      "json"
    ],
    "json_omitempty": false,
    "exclude_tables": [],
    "try_complete": true,
    "tinyint2bool": true,
    "sourceMap": [],
    "reflect": "{\"tinyint\":\"int8\",\"smallint\":\"int16\",\"int\":\"int32\",\"bigint\":\"int64\",\"float\":\"float64\",\"double\":\"float64\",\"decimal\":\"float64\",\"char\":\"string\",\"varchar\":\"string\",\"text\":\"string\",\"mediumtext\":\"string\",\"longtext\":\"string\",\"time\":\"time.Time\",\"date\":\"time.Time\",\"datetime\":\"time.Time\",\"timestamp\":\"time.Time\",\"enum\":\"string\",\"set\":\"string\",\"blob\":\"string\"}",
    "special": "{\"id\":\"uint\"}"
  }
}`)
	CustomFile = "./config.json"
)

func init() {
	_, err := os.Stat(CustomFile)
	if err == nil {
		Json, _ = ioutil.ReadFile(CustomFile)
	}
}
