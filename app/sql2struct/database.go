/*
@Time : 2019/12/20 16:06
@Software: GoLand
@File : database_tab
@Author : Bingo <airplayx@gmail.com>
*/
package sql2struct

import (
	"errors"
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
	"gormat/common"
	"strings"
	"time"
)

func DataBase(win fyne.Window, options *common.SQL2Struct) fyne.Widget {
	driver := widget.NewSelect([]string{"Mysql" /*, "PostgreSQL", "Sqlite3", "Mssql"*/}, func(s string) {

	})
	driver.SetSelected(strings.Title(options.Driver))
	host := widget.NewEntry()
	host.SetPlaceHolder("localhost")
	host.SetText(options.SourceMap.Host)
	port := widget.NewEntry()
	port.SetPlaceHolder("3306")
	port.SetText(options.SourceMap.Port)
	password := widget.NewPasswordEntry()
	password.SetPlaceHolder("Password")
	password.SetText(options.SourceMap.Password)
	user := widget.NewEntry()
	user.SetPlaceHolder("root")
	user.SetText(options.SourceMap.User)
	database := widget.NewEntry()
	database.SetText(options.SourceMap.Db)
	testDb := widget.NewHBox()
	testDb.Append(widget.NewButton("测试连接", func() {
		progressDialog := dialog.NewProgress("连接中", host.Text, win)
		go func() {
			num := 0.0
			for num < 1.0 {
				time.Sleep(50 * time.Millisecond)
				progressDialog.SetValue(num)
				num += 0.01
			}
			progressDialog.SetValue(1)
			progressDialog.Hide()
		}()
		progressDialog.Show()
		if err := common.InitDb(); err != nil {
			dialog.ShowError(errors.New("连接失败"), win)
		} else {
			dialog.ShowInformation("成功", "连接成功", win)
		}
	}))
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
			{Text: "", Widget: testDb},
		},
	}
}
