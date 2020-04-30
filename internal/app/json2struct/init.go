/*
@Time : 2020/1/3 14:28
@Software: GoLand
@File : init
@Author : Bingo <airplayx@gmail.com>
*/
package json2struct

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/airplayx/json2go"
	"strings"
)

func Screen() *fyne.Container {
	result := widget.NewMultiLineEntry()
	data := widget.NewMultiLineEntry()
	data.SetPlaceHolder(`{"a":"b"}`)
	result.SetPlaceHolder(`type YourStruct struct {
    A string ` + "`" + `json:"a"` + "`" + ` // b
}`)
	data.OnChanged = func(s string) {
		if s == "" {
			result.SetText(s)
			return
		}
		bytes, err := json2go.New([]byte(s), "test")
		if err != nil {
			result.SetText(err.Error())
			return
		}
		if b, err := bytes.WriteGo(); err != nil {
			result.SetText(err.Error())
			return
		} else {
			result.SetText(strings.ReplaceAll(string(b), "\t", "    "))
		}
	}

	return fyne.NewContainerWithLayout(
		layout.NewGridLayoutWithRows(1),
		widget.NewScrollContainer(data),
		widget.NewScrollContainer(result),
	)
}
