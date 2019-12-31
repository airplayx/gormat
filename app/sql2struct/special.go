/*
@Time : 2019/12/20 16:36
@Software: GoLand
@File : special_tab
@Author : Bingo <airplayx@gmail.com>
*/
package sql2struct

import (
	"encoding/json"
	"errors"
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
	"github.com/buger/jsonparser"
	_app "gormat/app"
	"gormat/controllers/Sql2struct"
	"strings"
)

func Special(win fyne.Window, options *Sql2struct.SQL2Struct) fyne.Widget {
	specialData := widget.NewMultiLineEntry()
	specialData.SetText(strings.ReplaceAll(options.Special, ",", ",\n"))
	return &widget.Form{
		OnCancel: func() {

		},
		OnSubmit: func() {
			options.Special = strings.ReplaceAll(specialData.Text, ",\n", ",")
			jsons, _ := json.Marshal(options)
			if data, err := jsonparser.Set(_app.Config, jsons, "sql2struct"); err == nil {
				_app.Config = data
				dialog.ShowInformation("成功", "保存成功", win)
			} else {
				dialog.ShowError(errors.New(err.Error()), win)
			}
		},
		Items: []*widget.FormItem{
			{Text: "指定字段名转型", Widget: specialData},
		},
	}
}
