package Sql2struct

import (
	"encoding/json"
	"github.com/buger/jsonparser"
	_app "gormat/app"
	"log"
)

type SQL2Struct struct {
	TargetDir     string   `json:"target_dir"`
	AutoSave      bool     `json:"auto_save"`
	Tags          []string `json:"tags"`
	JSONOmitempty bool     `json:"json_omitempty"`
	Driver        string   `json:"driver"`
	ExcludeTables []string `json:"exclude_tables"`
	TryComplete   bool     `json:"try_complete"`
	SourceMap     struct {
		User     string `json:"user"`
		Password string `json:"password"`
		Host     string `json:"host"`
		Port     string `json:"port"`
		Db       string `json:"db"`
	} `json:"sourceMap"`
	Reflect struct {
		Tinyint    string `json:"tinyint"`
		Smallint   string `json:"smallint"`
		Int        string `json:"int"`
		Bigint     string `json:"bigint"`
		Float      string `json:"float"`
		Double     string `json:"double"`
		Decimal    string `json:"decimal"`
		Char       string `json:"char"`
		Varchar    string `json:"varchar"`
		Text       string `json:"text"`
		Mediumtext string `json:"mediumtext"`
		Longtext   string `json:"longtext"`
		Time       string `json:"time"`
		Date       string `json:"date"`
		Datetime   string `json:"datetime"`
		Timestamp  string `json:"timestamp"`
		Enum       string `json:"enum"`
		Set        string `json:"set"`
		Blob       string `json:"blob"`
	} `json:"reflect"`
	Special struct {
		ID string `json:"id"`
	} `json:"special"`
}

func Configs() (config SQL2Struct) {
	data, _, _, _ := jsonparser.Get(_app.Config, "sql2struct")
	if err := json.Unmarshal([]byte(data), &config); err != nil {
		log.Println(err.Error())
	}
	return
}
