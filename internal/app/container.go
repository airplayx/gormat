/*
@Time : 2019/12/23 10:24
@Software: GoLand
@File : container
@Author : Bingo <airplayx@gmail.com>
*/
package app

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"gormat/internal/app/sql2struct"
	"gormat/internal/pkg/icon"
	"gormat/pkg/Sql2struct"
	"strings"
)

func Container(app fyne.App, win fyne.Window) *widget.TabContainer {
	var options = Sql2struct.Configs()
	var ipBox = widget.NewTabContainer()
	for _, v := range options.SourceMap {
		var dbBox = widget.NewTabContainer()
		for _, curDb := range v.Db {
			dbBox.Append(widget.NewTabItemWithIcon(
				curDb, icon.Database,
				sql2struct.Screen(win, &Sql2struct.SourceMap{
					Driver:   v.Driver,
					Host:     v.Host,
					User:     v.User,
					Password: v.Password,
					Port:     v.Port,
					Db:       []string{curDb},
				})))
		}
		if len(dbBox.Items) != 0 {
			dbBox.SetTabLocation(widget.TabLocationLeading)
		}
		i := icon.Mysql
		switch strings.Title(v.Driver) {
		case "PostgreSQL":
			i = icon.PostGreSQL
		case "Sqlite3":
			i = icon.SqLite
		case "Mssql":
			i = icon.Mssql
		}
		ipBox.Append(widget.NewTabItemWithIcon(v.Host+":"+v.Port, i, dbBox))
	}
	if len(ipBox.Items) == 0 {
		ipBox.Hide()
	} else {
		ipBox.SetTabLocation(widget.TabLocationLeading)
	}
	toolBar := ToolBar(win, ipBox, options)
	s2sBox := fyne.NewContainerWithLayout(
		layout.NewBorderLayout(toolBar, nil, nil, nil),
		toolBar,
		WelcomeScreen(),
		ipBox,
	)
	c := widget.NewTabContainer(
		//widget.NewTabItemWithIcon("", icon.Home, WelcomeScreen()),
		widget.NewTabItemWithIcon("", icon.Store, s2sBox),
		//widget.NewTabItemWithIcon("", theme.SettingsIcon(), SettingScreen(app, win)),
		//widget.NewTabItemWithIcon("", icon.Daily, fyne.NewContainer()),
		//widget.NewTabItemWithIcon("", icon.Video, fyne.NewContainer()),
	)
	c.SetTabLocation(widget.TabLocationBottom)
	return c
}
