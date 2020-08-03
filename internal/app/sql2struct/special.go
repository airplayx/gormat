/*Package sql2struct ...
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
	"gormat/pkg/sql2struct"
	"strings"
)

//Special ...
func Special(win fyne.Window, options *sql2struct.SQL2Struct) fyne.Widget {
	specialData := widget.NewMultiLineEntry()
	specialData.SetText(strings.ReplaceAll(options.Special, ",", ",\n"))
	autoBool := widget.NewCheck(configs.Text("the fields starting with is_ are automatically converted to bool"), func(bool) {})
	autoBool.SetChecked(options.Tinyint2bool)

	return &widget.Form{
		OnCancel: func() {
			//win.Close()
		},
		OnSubmit: func() {
			options.Special = strings.ReplaceAll(specialData.Text, ",\n", ",")
			options.Tinyint2bool = autoBool.Checked
			jsons, _ := json.Marshal(options)
			if data, err := jsonparser.Set(configs.JSON, jsons, "sql2struct"); err == nil {
				configs.JSON = data
				dialog.ShowInformation(configs.Text("info"), configs.Text("save ok"), win)
			} else {
				dialog.ShowError(errors.New(err.Error()), win)
			}
		},
		Items: []*widget.FormItem{
			{Text: configs.Text("bool"), Widget: autoBool},
			{Text: configs.Text("specified fields transformation"), Widget: specialData},
		},
		CancelText: configs.Text("cancel"),
		SubmitText: configs.Text("confirm"),
	}
}
