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
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	_app "gormat/app"
	"gormat/app/json2struct"
	"gormat/app/sql2struct"
	"gormat/controllers/Sql2struct"
)

func Aside(app fyne.App, win fyne.Window) (aside *widget.TabContainer) {
	var options = Sql2struct.Configs()
	sql2Screen := sql2struct.Screen(win)
	aside = widget.NewTabContainer(
		widget.NewTabItemWithIcon("", theme.HomeIcon(), _app.WelcomeScreen()),
		//widget.NewTabItemWithIcon("", theme.SettingsIcon(), _app.SettingScreen(app, win)),
		widget.NewTabItem("Sql转Struct", fyne.NewContainerWithLayout(
			layout.NewGridLayout(1),
			widget.NewTabContainer(
				widget.NewTabItem("选项", sql2struct.Option(win, options)),
				widget.NewTabItem("映射", sql2struct.Reflect(win, options)),
				widget.NewTabItem("特殊转型", sql2struct.Special(win, options)),
				widget.NewTabItem("数据库", sql2struct.DataBase(win, sql2Screen, options)),
				widget.NewTabItemWithIcon("开始转换", theme.ViewRefreshIcon(), sql2Screen),
			),
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
