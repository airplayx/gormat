/*
@Time : 2019/12/20 16:36
@Software: GoLand
@File : special
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
	"gormat/configs"
	"gormat/pkg/Sql2struct"
	"strings"
)

func Special(win fyne.Window, options *Sql2struct.SQL2Struct) fyne.Widget {
	specialData := widget.NewMultiLineEntry()
	specialData.SetText(strings.ReplaceAll(options.Special, ",", ",\n"))
	autoBool := widget.NewCheck("is_开头的字段自动转bool", func(bool) {})
	autoBool.SetChecked(options.Tinyint2bool)

	return &widget.Form{
		OnCancel: func() {
			win.Close()
		},
		OnSubmit: func() {
			options.Special = strings.ReplaceAll(specialData.Text, ",\n", ",")
			options.Tinyint2bool = autoBool.Checked
			jsons, _ := json.Marshal(options)
			if data, err := jsonparser.Set(configs.Json, jsons, "sql2struct"); err == nil {
				configs.Json = data
				dialog.ShowInformation("成功", "保存成功", win)
			} else {
				dialog.ShowError(errors.New(err.Error()), win)
			}
		},
		Items: []*widget.FormItem{
			{Text: "布尔选型", Widget: autoBool},
			{Text: "指定字段转型", Widget: specialData},
		},
	}
}
