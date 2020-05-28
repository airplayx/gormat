/*Package sql2struct ...
@Time : 2019/12/20 16:06
@Software: GoLand
@File : database
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
	"gormat/internal/pkg/icon"
	"gormat/pkg/sql2struct"
	"strings"
	"time"
)

//DataBase create new dbConf or load
func DataBase(window, currentWindow fyne.Window, ipBox *widget.TabContainer, options *sql2struct.SQL2Struct, dbIndex []int) fyne.Widget {
	driver := widget.NewSelect([]string{"Mysql" /*, "PostgreSQL", "Sqlite3", "Mssql"*/}, func(s string) {

	})
	host := widget.NewEntry()
	host.SetPlaceHolder("localhost")
	port := widget.NewEntry()
	port.SetPlaceHolder("3306")
	password := widget.NewPasswordEntry()
	password.SetPlaceHolder("******")
	user := widget.NewEntry()
	user.SetPlaceHolder("root")
	db := widget.NewEntry()
	driver.SetSelected("Mysql")
	if dbIndex != nil {
		currentLink := options.SourceMap[dbIndex[0]]
		driver.SetSelected(strings.Title(currentLink.Driver))
		host.SetText(currentLink.Host)
		port.SetText(currentLink.Port)
		password.SetText(currentLink.Password)
		user.SetText(currentLink.User)
		if len(dbIndex) > 1 {
			db.SetText(currentLink.Db[dbIndex[1]])
		}
	}
	testDb := widget.NewHBox(widget.NewButton(configs.Text("connection test"), func() {
		progressDialog := dialog.NewProgress(configs.Text("testing"), host.Text, currentWindow)
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
		err := sql2struct.Init(&sql2struct.SourceMap{
			Db:       []string{db.Text},
			User:     user.Text,
			Password: password.Text,
			Host:     host.Text,
			Port:     port.Text,
			Driver:   driver.Selected,
		})
		if err != nil {
			dialog.ShowError(errors.New(err.Error()), currentWindow)
		} else {
			dialog.ShowInformation(configs.Text("info"), configs.Text("connection successful"), currentWindow)
			_ = sql2struct.Engine.Close()
		}
	}))
	return &widget.Form{
		OnCancel: func() {
			currentWindow.Close()
		},
		OnSubmit: func() {
			if db.Text == "" {
				dialog.ShowError(errors.New(configs.Text("the database can not be empty")), currentWindow)
				return
			}
			newDB := widget.NewTabItemWithIcon(
				db.Text, icon.Database,
				Screen(currentWindow, &sql2struct.SourceMap{
					Driver:   driver.Selected,
					Host:     host.Text,
					User:     user.Text,
					Password: password.Text,
					Port:     port.Text,
					Db:       []string{db.Text},
				}))
			var dbBox = widget.NewTabContainer(newDB)
			dbBox.SetTabLocation(widget.TabLocationLeading)
			defer func() {
				window.Canvas().Refresh(ipBox)
				dbBox.SelectTab(newDB)
			}()
			i := icon.Mysql
			switch strings.Title(driver.Selected) {
			case "PostgreSQL":
				i = icon.PostGreSQL
			case "Sqlite3":
				i = icon.SqLite
			case "Mssql":
				i = icon.Mssql
			}
			sourceMap := options.SourceMap
			oldIP := false
			for _, v := range sourceMap {
				if v.Host+":"+v.Port == host.Text+":"+port.Text {
					for _, curDb := range v.Db {
						if curDb == db.Text && v.Password == password.Text {
							dialog.ShowError(errors.New(configs.Text("duplicate links")), currentWindow)
							return
						}
					}
					oldIP = true
				}
			}
			if oldIP {
				dbBox = ipBox.CurrentTab().Content.(*widget.TabContainer)
				if dbIndex != nil && len(dbIndex) > 1 {
					currentLink := sourceMap[dbIndex[0]]
					currentLink.Driver = driver.Selected
					currentLink.Db[dbIndex[1]] = db.Text
					currentLink.User = user.Text
					currentLink.Password = password.Text
					currentLink.Host = host.Text
					currentLink.Port = port.Text
					dbBox.RemoveIndex(dbBox.CurrentTabIndex())
				}
				for k, v := range sourceMap {
					if v.Host+":"+v.Port == host.Text+":"+port.Text {
						ipBox.SelectTabIndex(k)
						ipBox.CurrentTab().Content.(*widget.TabContainer).Append(newDB)
						sourceMap[k].Db = sql2struct.RmDuplicateElement(append(v.Db, db.Text))
						break
					}
				}
			} else {
				ipBox.Append(widget.NewTabItemWithIcon(host.Text+":"+port.Text, i, dbBox))
				ipBox.SetTabLocation(widget.TabLocationLeading)
				ipBox.SelectTabIndex(len(ipBox.Items) - 1)
				options.SourceMap = append(sourceMap, sql2struct.SourceMap{
					Driver:   driver.Selected,
					Host:     host.Text,
					User:     user.Text,
					Password: password.Text,
					Port:     port.Text,
					Db:       []string{db.Text},
				})
			}
			jsons, _ := json.Marshal(options)
			if data, err := jsonparser.Set(configs.JSON, jsons, "sql2struct"); err == nil {
				configs.JSON = data
				defer func() {
					if ipBox.Hidden {
						ipBox.Show()
					}
					currentWindow.Close()
				}()
			} else {
				dialog.ShowError(errors.New(err.Error()), currentWindow)
			}
		},
		Items: []*widget.FormItem{
			{Text: configs.Text("engine"), Widget: driver},
			{Text: configs.Text("host"), Widget: host},
			{Text: configs.Text("port"), Widget: port},
			{Text: configs.Text("user"), Widget: user},
			{Text: configs.Text("password"), Widget: password},
			{Text: configs.Text("database"), Widget: db},
			{Text: "", Widget: testDb},
		},
	}
}
