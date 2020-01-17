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
	Json       = []byte(``)
	CustomFile = "./config.json"
)

func init() {
	Json, _ = ioutil.ReadFile(`./configs/default.json`)
	_, err := os.Stat(CustomFile)
	if err == nil {
		Json, _ = ioutil.ReadFile(CustomFile)
	}
}
