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
	"gormat/app/config"
	"gormat/internal/Sql2struct"
	"strings"
	"time"
	"xorm.io/core"
	"xorm.io/xorm"
)

func DataBase(win fyne.Window, options *Sql2struct.SQL2Struct, dbIndex int) fyne.Widget {
	driver := widget.NewSelect([]string{"Mysql" /*, "PostgreSQL", "Sqlite3", "Mssql"*/}, func(s string) {

	})
	driver.SetSelected(strings.Title(options.Driver))
	host := widget.NewEntry()
	host.SetPlaceHolder("localhost")
	port := widget.NewEntry()
	port.SetPlaceHolder("3306")
	password := widget.NewPasswordEntry()
	password.SetPlaceHolder("Password")
	user := widget.NewEntry()
	user.SetPlaceHolder("root")
	db := widget.NewEntry()
	if dbIndex > -1 {
		host.SetText(options.SourceMap[dbIndex].Host)
		port.SetText(options.SourceMap[dbIndex].Port)
		password.SetText(options.SourceMap[dbIndex].Password)
		user.SetText(options.SourceMap[dbIndex].User)
		db.SetText(options.SourceMap[dbIndex].Db[0])
	}
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
				db.Text,
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
	return &widget.Form{
		OnCancel: func() {
			win.Close()
		},
		OnSubmit: func() {
			options.Driver = driver.Selected
			if dbIndex > -1 {
				//options.SourceMap[dbIndex].Db[0] = db.Text
				options.SourceMap[dbIndex].User = user.Text
				options.SourceMap[dbIndex].Password = password.Text
				options.SourceMap[dbIndex].Host = host.Text
				options.SourceMap[dbIndex].Port = port.Text
			}
			jsons, _ := json.Marshal(options)
			if data, err := jsonparser.Set(config.Setting, jsons, "sql2struct"); err == nil {
				config.Setting = data
				dialog.ShowInformation("成功", "保存成功", win)
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
			{Text: "数据库", Widget: db},
			{Text: "", Widget: testDb},
		},
	}
}
