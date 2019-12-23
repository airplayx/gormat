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
)

func Aside(win fyne.Window) (aside *widget.TabContainer) {
	sql2Str := widget.NewTabContainer(
		widget.NewTabItem("选项", sql2struct.Option()),
		widget.NewTabItem("数据库", sql2struct.DataBase()),
		widget.NewTabItem("映射", sql2struct.Reflect(win)),
		widget.NewTabItem("特殊转型", sql2struct.Special(win)))
	if c := _app.Sql2structScreen(win); len(c.Objects) != 0 {
		sql2Str.Items = append(sql2Str.Items,
			widget.NewTabItemWithIcon("开始转换", theme.ViewRefreshIcon(), c),
		)
	}
	aside = widget.NewTabContainer(
		widget.NewTabItemWithIcon("首页", nil, _app.WelcomeScreen()),
		widget.NewTabItemWithIcon("SQL转Struct", nil, fyne.NewContainerWithLayout(
			layout.NewGridLayoutWithColumns(1),
			sql2Str,
		)),
		widget.NewTabItemWithIcon("JSON转Struct", nil, fyne.NewContainerWithLayout(
			layout.NewGridLayoutWithColumns(1),
		)),
		widget.NewTabItemWithIcon("URL编解码", nil, fyne.NewContainerWithLayout(
			layout.NewGridLayoutWithColumns(1),
		)),
		widget.NewTabItemWithIcon("时间戳格式", nil, fyne.NewContainerWithLayout(
			layout.NewGridLayoutWithColumns(1),
		)),
	)
	aside.SetTabLocation(widget.TabLocationLeading)
	return
}
