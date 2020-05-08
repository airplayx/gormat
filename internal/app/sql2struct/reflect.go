/*Package sql2struct ...
@Time : 2019/12/20 16:36
@Software: GoLand
@File : reflect
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
	"gormat/pkg/sql2struct"
	"strings"
)

//Reflect ...
func Reflect(win fyne.Window, options *sql2struct.SQL2Struct) fyne.Widget {
	dataType := widget.NewMultiLineEntry()
	dataType.SetText(strings.ReplaceAll(options.Reflect, ",", ",\n"))
	return &widget.Form{
		OnCancel: func() {
			win.Close()
		},
		OnSubmit: func() {
			options.Reflect = strings.ReplaceAll(dataType.Text, ",\n", ",")
			jsons, _ := json.Marshal(options)
			if data, err := jsonparser.Set(configs.JSON, jsons, "sql2struct"); err == nil {
				configs.JSON = data
				dialog.ShowInformation(configs.Text("info"), configs.Text("save ok"), win)
			} else {
				dialog.ShowError(errors.New(err.Error()), win)
			}
		},
		Items: []*widget.FormItem{
			{Text: configs.Text("type conversion"), Widget: dataType},
		},
	}
}
