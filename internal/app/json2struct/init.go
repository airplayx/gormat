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
	"gormat/pkg/Json2struct"
	"strings"
)

func Screen() *fyne.Container {
	result := widget.NewMultiLineEntry()
	data := widget.NewMultiLineEntry()
	data.OnChanged = func(s string) {
		if s == "" {
			result.SetText(s)
			return
		}
		f, err := Json2struct.ParseJson([]byte(s))
		if err != nil {
			result.SetText(err.Error())
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
