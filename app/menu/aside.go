/*
@Time : 2019/12/23 10:24
@Software: GoLand
@File : aside
@Author : Bingo <airplayx@gmail.com>
*/
package menu

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	_app "gormat/app"
	"gormat/app/json2struct"
	"gormat/app/sql2struct"
	"gormat/controllers/Sql2struct"
)

func Aside(app fyne.App, win fyne.Window) (aside *widget.TabContainer) {
	var options = Sql2struct.Configs()
	var ipList []*widget.TabItem
	for _, v := range options.SourceMap {
		var dbList []*widget.TabItem
		for _, curDb := range v.Db {
			dbList = append(dbList, widget.NewTabItemWithIcon(curDb, _app.Database,
				sql2struct.Screen(win, []interface{}{
					v.User,
					v.Password,
					v.Host,
					v.Port,
					curDb,
				})))
		}
		dbBox := widget.NewTabContainer(dbList...)
		dbBox.SetTabLocation(widget.TabLocationLeading)
		ipList = append(ipList, widget.NewTabItem(v.Host, dbBox))
	}
	IPBox := widget.NewTabContainer(ipList...)
	addBox := widget.NewScrollContainer(
		sql2struct.DataBase(win, sql2struct.Screen(win, []interface{}{}), options, 0),
	)
	addBox.Hide()

	setting := widget.NewTabContainer(
		widget.NewTabItem("选项", sql2struct.Option(win, options)),
		widget.NewTabItem("映射", sql2struct.Reflect(win, options)),
		widget.NewTabItem("特殊转型", sql2struct.Special(win, options)),
	)
	setting.SetTabLocation(widget.TabLocationLeading)
	setting.Hide()
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(_app.Line, func() {
			IPBox.Show()
			setting.Hide()
			addBox.Hide()
		}),
		//widget.NewToolbarSeparator(),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(_app.Insert, func() {
			IPBox.Hide()
			setting.Hide()
			addBox.Show()
		}),
		widget.NewToolbarAction(_app.Setting, func() {
			IPBox.Hide()
			setting.Show()
			addBox.Hide()
		}),
	)
	s2sBox := fyne.NewContainerWithLayout(
		layout.NewBorderLayout(toolbar, nil, nil, nil),
		toolbar,
		addBox,
		setting,
	)
	if len(IPBox.Items) > 0 {
		IPBox.SetTabLocation(widget.TabLocationLeading)
		s2sBox.AddObject(IPBox)
	}
	aside = widget.NewTabContainer(
		widget.NewTabItemWithIcon("", _app.Home, _app.WelcomeScreen()),
		//widget.NewTabItemWithIcon("", theme.SettingsIcon(), _app.SettingScreen(app, win)),
		widget.NewTabItem("Sql转Struct", s2sBox),
		widget.NewTabItem("Json转Struct", fyne.NewContainerWithLayout(
			layout.NewGridLayout(1),
			widget.NewScrollContainer(json2struct.Screen()),
		)),
		//widget.NewTabItem("日期格式化", fyne.NewContainerWithLayout(
		//	layout.NewGridLayout(1),
		//)),
		//widget.NewTabItem("URL编解码", fyne.NewContainerWithLayout(
		//	layout.NewGridLayout(1),
		//)),
	)
	aside.SetTabLocation(widget.TabLocationBottom)
	return
}
