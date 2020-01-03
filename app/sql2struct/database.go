/*
@Time : 2019/12/20 16:06
@Software: GoLand
@File : database_tab
@Author : Bingo <airplayx@gmail.com>
*/
package sql2struct

import (
	"encoding/json"
	"errors"
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
	"github.com/buger/jsonparser"
	_app "gormat/app"
	"gormat/controllers/Sql2struct"
	"strings"
	"time"
	"xorm.io/core"
	"xorm.io/xorm"
)

func DataBase(win fyne.Window, tab *fyne.Container, options *Sql2struct.SQL2Struct) fyne.Widget {
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
	testDb := widget.NewHBox(widget.NewButton("测试连接", func() {
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
		engine, err := xorm.NewEngine(
			strings.ToLower(driver.Selected),
			fmt.Sprintf("%s:%s@(%s:%s)/%s",
				user.Text,
				password.Text,
				host.Text,
				port.Text,
				database.Text,
			))
		if err != nil {
			dialog.ShowError(errors.New(err.Error()), win)
			return
		}
		engine.SetLogLevel(core.LOG_WARNING)
		if err := engine.Ping(); err != nil {
			dialog.ShowError(errors.New(err.Error()), win)
		} else {
			dialog.ShowInformation("成功", "连接成功", win)
		}
		_ = engine.Close()
	}))
	go func() {
		_ = initTables(win, tab)
	}()
	return &widget.Form{
		OnCancel: func() {

		},
		OnSubmit: func() {
			options.Driver = driver.Selected
			options.SourceMap.Db = database.Text
			options.SourceMap.User = user.Text
			options.SourceMap.Password = password.Text
			options.SourceMap.Host = host.Text
			options.SourceMap.Port = port.Text
			jsons, _ := json.Marshal(options)
			if data, err := jsonparser.Set(_app.Config, jsons, "sql2struct"); err == nil {
				_app.Config = data
				if err := initTables(win, tab); err != nil {
					dialog.ShowError(errors.New(err.Error()), win)
				} else {
					dialog.ShowInformation("成功", "保存成功", win)
				}
			} else {
				dialog.ShowError(errors.New(err.Error()), win)
			}
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

func initTables(win fyne.Window, tab *fyne.Container) (err error) {
	if err = Sql2struct.InitDb(); err != nil {
		return
	}
	if Sql2struct.Tables, err = Sql2struct.DBMetas(nil,
		Sql2struct.Configs().ExcludeTables,
		Sql2struct.Configs().TryComplete); err == nil {
		tab.Objects = Screen(win).Objects
	}
	return
}
