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
	"gormat/app/sql2struct"
	"gormat/controllers/Sql2struct"
)

func Aside(app fyne.App, win fyne.Window) (aside *widget.TabContainer) {
	var options = Sql2struct.Configs()
	sql2Screen := sql2struct.Screen(win)
	aside = widget.NewTabContainer(
		widget.NewTabItemWithIcon("首页", nil, _app.WelcomeScreen()),
		widget.NewTabItemWithIcon("Sql转Struct", nil, fyne.NewContainerWithLayout(
			layout.NewGridLayoutWithColumns(1),
			widget.NewTabContainer(
				widget.NewTabItem("选项", sql2struct.Option(win, options)),
				widget.NewTabItem("映射", sql2struct.Reflect(win, options)),
				widget.NewTabItem("特殊转型", sql2struct.Special(win, options)),
				widget.NewTabItem("数据库", sql2struct.DataBase(win, sql2Screen, options)),
				widget.NewTabItemWithIcon("开始转换", theme.ViewRefreshIcon(), sql2Screen),
			),
		)),
		widget.NewTabItemWithIcon("Json转Struct", nil, fyne.NewContainerWithLayout(
			layout.NewGridLayoutWithColumns(1),
		)),
		widget.NewTabItemWithIcon("Map转Struct", nil, fyne.NewContainerWithLayout(
			layout.NewGridLayoutWithColumns(1),
		)),
		widget.NewTabItemWithIcon("日期格式化", nil, fyne.NewContainerWithLayout(
			layout.NewGridLayoutWithColumns(1),
		)),
		widget.NewTabItemWithIcon("URL编解码", nil, fyne.NewContainerWithLayout(
			layout.NewGridLayoutWithColumns(1),
		)),
		//widget.NewTabItemWithIcon("设置", nil, _app.SettingScreen(app, win)),
	)
	aside.SetTabLocation(widget.TabLocationBottom)
	return
}
