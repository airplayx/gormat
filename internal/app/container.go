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
)

func Container(app fyne.App, win fyne.Window) *widget.TabContainer {
	var options = Sql2struct.Configs()
	var dbBox, ipBox = widget.NewTabContainer(), widget.NewTabContainer()
	for _, v := range options.SourceMap {
		for _, curDb := range v.Db {
			dbBox.Items = append(dbBox.Items, widget.NewTabItemWithIcon(
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
		dbBox.SetTabLocation(widget.TabLocationLeading)
		ipBox.Items = append(ipBox.Items, widget.NewTabItemWithIcon(v.Host+":"+v.Port, icon.Mysql, dbBox))
	}
	toolBar := ToolBar(win, ipBox, dbBox, options)
	s2sBox := fyne.NewContainerWithLayout(
		layout.NewBorderLayout(toolBar, nil, nil, nil),
		toolBar,
		WelcomeScreen(),
	)
	if len(ipBox.Items) > 0 {
		ipBox.SetTabLocation(widget.TabLocationLeading)
		s2sBox.AddObject(ipBox)
	}
	c := widget.NewTabContainer(
		//widget.NewTabItemWithIcon("", config.Home, WelcomeScreen()),
		//widget.NewTabItemWithIcon("", theme.SettingsIcon(), _app.SettingScreen(app, win)),
		widget.NewTabItemWithIcon("", icon.Store, s2sBox),
		//widget.NewTabItemWithIcon("", config.Daily, fyne.NewContainer()),
		//widget.NewTabItemWithIcon("", config.Video, fyne.NewContainer()),
	)
	c.SetTabLocation(widget.TabLocationBottom)
	return c
}
