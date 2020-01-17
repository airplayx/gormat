/*
@Time : 2019/12/23 10:24
@Software: GoLand
@File : aside
@Author : Bingo <airplayx@gmail.com>
*/
package _app

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"gormat/app/config"
	"gormat/app/sql2struct"
	"gormat/internal/Sql2struct"
	"time"
)

func Container(app fyne.App, win fyne.Window) *widget.TabContainer {
	var options = Sql2struct.Configs()
	var dbBox, ipBox = widget.NewTabContainer(), widget.NewTabContainer()
	currentIP, currentDB := make(chan *widget.TabItem), make(chan *widget.TabItem)
	for _, v := range options.SourceMap {
		for _, curDb := range v.Db {
			dbBox.Items = append(dbBox.Items, widget.NewTabItemWithIcon(
				curDb, config.Database,
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
		ipBox.Items = append(ipBox.Items, widget.NewTabItem(v.Host, dbBox))
	}
	toolBar := ToolBar(win, ipBox, dbBox, options)
	s2sBox := fyne.NewContainerWithLayout(
		layout.NewBorderLayout(toolBar, nil, nil, nil),
		toolBar,
		WelcomeScreen(),
	)
	go func() {
		for {
			time.Sleep(time.Microsecond * 200)
			if <-currentDB != dbBox.CurrentTab() {

			}
			if <-currentIP != ipBox.CurrentTab() {

			}
		}
	}()
	if len(ipBox.Items) > 0 {
		go func() {
			currentIP <- &widget.TabItem{}
			currentDB <- &widget.TabItem{}
			for {
				currentIP <- ipBox.CurrentTab()
				currentDB <- dbBox.CurrentTab()
			}
		}()
		ipBox.SetTabLocation(widget.TabLocationLeading)
		s2sBox.AddObject(ipBox)
	}
	c := widget.NewTabContainer(
		//widget.NewTabItemWithIcon("", config.Home, WelcomeScreen()),
		//widget.NewTabItemWithIcon("", theme.SettingsIcon(), _app.SettingScreen(app, win)),
		widget.NewTabItemWithIcon("", config.Store, s2sBox),
		widget.NewTabItemWithIcon("", config.Daily, fyne.NewContainer()),
		//widget.NewTabItemWithIcon("", config.Video, fyne.NewContainer()),
	)
	c.SetTabLocation(widget.TabLocationBottom)
	return c
}
