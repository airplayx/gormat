/*
@Time : 2019/12/20 16:36
@Software: GoLand
@File : reflect_tab
@Author : Bingo <airplayx@gmail.com>
*/
package sql2struct

import (
	"encoding/json"
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
	"gormat/controllers/Sql2struct"
	"strings"
	"time"
)

func Reflect(win fyne.Window, options *Sql2struct.SQL2Struct) fyne.Widget {
	dataType := widget.NewMultiLineEntry()
	reflect, _ := json.Marshal(options.Reflect)
	dataType.SetText(strings.ReplaceAll(string(reflect), ",", ",\n"))
	return &widget.Form{
		OnCancel: func() {
			cnf := dialog.NewConfirm("Confirmation", "Are you enjoying this demo?", func(b bool) {
				fmt.Println("Responded with", b)
			}, win)
			cnf.SetDismissText("Nah")
			cnf.SetConfirmText("Oh Yes!")
			cnf.Show()
		},
		OnSubmit: func() {
			prog := dialog.NewProgress("MyProgress", "Nearly there...", win)

			go func() {
				num := 0.0
				for num < 1.0 {
					time.Sleep(50 * time.Millisecond)
					prog.SetValue(num)
					num += 0.01
				}

				prog.SetValue(1)
				prog.Hide()
			}()

			prog.Show()
		},
		Items: []*widget.FormItem{
			{Text: "数据类型转换", Widget: dataType},
		},
	}
}
