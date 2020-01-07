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
	sql2Screen := sql2struct.Screen(win)
	moreBox := widget.NewTabContainer(
		widget.NewTabItemWithIcon("test", _app.Database, sql2Screen),
		widget.NewTabItemWithIcon("test2", _app.Database, fyne.NewContainer()),
	)
	addBox := widget.NewScrollContainer(
		sql2struct.DataBase(win, sql2Screen, options),
	)
	addBox.Hide()
	moreBox.SetTabLocation(widget.TabLocationLeading)
	dataBox := widget.NewTabContainer(
		widget.NewTabItem("localhost", fyne.NewContainer()),
		widget.NewTabItem("127.0.0.1", moreBox),
	)
	dataBox.SetTabLocation(widget.TabLocationLeading)
	setting := widget.NewTabContainer(
		widget.NewTabItem("选项", sql2struct.Option(win, options)),
		widget.NewTabItem("映射", sql2struct.Reflect(win, options)),
		widget.NewTabItem("特殊转型", sql2struct.Special(win, options)),
	)
	setting.SetTabLocation(widget.TabLocationLeading)
	setting.Hide()
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(_app.Line, func() {
			dataBox.Show()
			setting.Hide()
		}),
		//widget.NewToolbarSeparator(),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(_app.Insert, func() {
			dataBox.Hide()
			setting.Hide()
			addBox.Show()
		}),
		widget.NewToolbarAction(_app.Setting, func() {
			dataBox.Hide()
			setting.Show()
		}),
	)
	aside = widget.NewTabContainer(
		widget.NewTabItemWithIcon("", _app.Home, _app.WelcomeScreen()),
		//widget.NewTabItemWithIcon("", theme.SettingsIcon(), _app.SettingScreen(app, win)),
		widget.NewTabItem("Sql转Struct", fyne.NewContainerWithLayout(
			layout.NewBorderLayout(toolbar, nil, nil, nil),
			toolbar,
			addBox,
			dataBox,
			setting,
		)),
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
