/*
@Time : 2019/12/20 16:06
@Software: GoLand
@File : database_tab
@Author : Bingo <airplayx@gmail.com>
*/
package _app

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

func DataBaseTab() fyne.Widget {
	host := widget.NewEntry()
	host.SetPlaceHolder("localhost")
	port := widget.NewEntry()
	port.SetPlaceHolder("3306")
	driver := widget.NewEntry()
	driver.SetPlaceHolder("mysql")
	driver.Disable()
	password := widget.NewPasswordEntry()
	password.SetPlaceHolder("Password")
	user := widget.NewEntry()
	user.SetPlaceHolder("root")
	database := widget.NewEntry()
	return &widget.Form{
		OnCancel: func() {

		},
		OnSubmit: func() {

		},
		Items: []*widget.FormItem{
			{Text: "host", Widget: host},
			{Text: "port", Widget: port},
			{Text: "driver", Widget: driver},
			{Text: "user", Widget: user},
			{Text: "password", Widget: password},
			{Text: "database", Widget: database},
		},
	}
}
