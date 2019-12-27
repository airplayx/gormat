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
	"gormat/common"
	"strings"
)

func Special(win fyne.Window, options *common.SQL2Struct) fyne.Widget {
	specialData := widget.NewMultiLineEntry()
	special, _ := json.Marshal(options.Special)
	specialData.SetText(strings.ReplaceAll(string(special), ",", ",\n"))
	return &widget.Form{
		OnCancel: func() {
			err := errors.New("A dummy error message")
			dialog.ShowError(err, win)
		},
		OnSubmit: func() {
			dialog.ShowInformation("Information", "You should know this thing...", win)
		},
		Items: []*widget.FormItem{
			{Text: "指定字段名转型", Widget: specialData},
		},
	}
}
