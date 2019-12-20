/*
@Time : 2019/12/20 16:40
@Software: GoLand
@File : option
@Author : Bingo <airplayx@gmail.com>
*/
package _app

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

func OptionTab() fyne.Widget {
	targetDir := widget.NewEntry()
	targetDir.SetPlaceHolder("./models")
	gorm := widget.NewCheck("gorm", func(bool) {})
	gorm.SetChecked(true)
	gorm.Hide()
	json := widget.NewCheck("yes", func(bool) {})
	jsonOmitempty := widget.NewRadio([]string{"omitempty", "no omitempty"}, func(s string) {

	})
	jsonOmitempty.Horizontal = true

	excludeTables := widget.NewMultiLineEntry()

	tryComplete := widget.NewRadio([]string{"yes", "no"}, func(s string) {

	})
	tryComplete.Horizontal = true
	return &widget.Form{
		OnCancel: func() {

		},
		OnSubmit: func() {

		},
		Items: []*widget.FormItem{
			{Text: "target dir", Widget: targetDir},
			{Text: "json", Widget: json},
			{Text: "", Widget: jsonOmitempty},
			{Text: "exclude tables", Widget: excludeTables},
			{Text: "try complete", Widget: tryComplete},
		},
	}
}
