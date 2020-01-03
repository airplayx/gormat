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
	"gormat/controllers/Json2struct"
	"strings"
)

func Screen() *fyne.Container {
	result := widget.NewMultiLineEntry()
	data := widget.NewMultiLineEntry()
	data.OnChanged = func(s string) {
		if data.Text == "" {
			result.SetText(data.Text)
			return
		}
		f, err := Json2struct.ParseJson([]byte(data.Text))
		if err != nil {
			result.SetText(err.Error())
			return
		}
		switch f.(type) {
		case map[string]interface{}:
		default:
			result.SetText("错误的数据")
			return
		}
		bytes, _ := Json2struct.PrintGo(f, "YourStruct")
		result.SetText(strings.ReplaceAll(string(bytes), "\t", "    "))
	}
	data.SetPlaceHolder(`{"a":"b"}`)
	result.SetPlaceHolder(`type YourStruct struct {
    A string ` + "`" + `json:"a"` + "`" + ` // b
}`)
	return fyne.NewContainerWithLayout(
		layout.NewGridLayoutWithRows(1),
		widget.NewScrollContainer(data),
		widget.NewScrollContainer(result),
	)
}
