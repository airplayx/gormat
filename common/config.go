package common

import (
	"encoding/json"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var (
	config      AppConfigs
	workPath, _ = os.Getwd()
)

type AppConfigs struct {
	TargetDir     string            `toml:"target_dir"`
	Driver        string            `toml:"driver"`
	Source        string            `toml:"source"`
	TagType       []string          `toml:"tag_type"`
	JsonOmitempty bool              `toml:"json_omitempty"`
	ExcludeTables []string          `toml:"exclude_tables"`
	TryComplete   bool              `toml:"try_complete"`
	Reflect       map[string]string `toml:"reflect"`
	Special       map[string]string `toml:"special"`
}

func LoadConfig() (err error) {
	appConfigPath := filepath.Join(workPath, "config.toml")
	if AppPath, err := filepath.Abs(filepath.Dir(os.Args[0])); err != nil {
		panic(err)
	} else {
		envConfig := filepath.Join(AppPath, "config.toml")
		if _, err := os.Stat(envConfig); err == nil {
			appConfigPath = filepath.Join(envConfig)
		}
	}
	_, err = toml.DecodeFile(appConfigPath, &config)
	return
}

func Configs() AppConfigs {
	reflect, err := ioutil.ReadFile("reflect.json")
	if err != nil {
		log.Fatalln(err.Error())
	}
	if err := json.Unmarshal(reflect, &config.Reflect); err != nil {
		log.Fatalln(err.Error())
	}
	special, err := ioutil.ReadFile("special.json")
	if err != nil {
		log.Fatalln(err.Error())
	}
	if err := json.Unmarshal(special, &config.Special); err != nil {
		log.Fatalln(err.Error())
	}
	return config
}
