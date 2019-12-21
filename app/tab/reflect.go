/*
@Time : 2019/12/20 16:36
@Software: GoLand
@File : reflect_tab
@Author : Bingo <airplayx@gmail.com>
*/
package tab

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
	"time"
)

func Reflect(win fyne.Window) fyne.Widget {
	dataType := widget.NewMultiLineEntry()
	dataType.SetText(`{
"tinyint": "int8",
"smallint": "int16",
"int": "int32",
"bigint": "int64",
"float": "float64",
"double": "float64",
"decimal": "float64",
"char": "string",
"varchar": "string",
"text": "string",
"mediumtext": "string",
"longtext": "string",
"time": "time.Time",
"date": "time.Time",
"datetime": "time.Time",
"timestramp": "int64",
"enum": "string",
"set": "string",
"blob": "string"
}`)
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
