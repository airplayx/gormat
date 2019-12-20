/*
@Time : 2019/12/20 16:06
@Software: GoLand
@File : database_tab
@Author : Bingo <airplayx@gmail.com>
*/
package tab

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

func DataBase() fyne.Widget {
	driver := widget.NewSelect([]string{"Mysql" /*, "PostgreSQL", "Sqlite3", "Mssql"*/}, func(s string) {

	})
	driver.SetSelected("Mysql")
	host := widget.NewEntry()
	host.SetPlaceHolder("localhost")
	port := widget.NewEntry()
	port.SetPlaceHolder("3306")
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
			{Text: "引擎", Widget: driver},
			{Text: "主机地址", Widget: host},
			{Text: "端口", Widget: port},
			{Text: "用户名", Widget: user},
			{Text: "密码", Widget: password},
			{Text: "数据库", Widget: database},
		},
	}
}
