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
	"gormat/common"
)

func Aside(win fyne.Window) (aside *widget.TabContainer) {
	var options = common.Configs()
	sql2Str := widget.NewTabContainer(
		widget.NewTabItem("选项", sql2struct.Option(&options)),
		widget.NewTabItem("数据库", sql2struct.DataBase()),
		widget.NewTabItem("映射", sql2struct.Reflect(win)),
		widget.NewTabItem("特殊转型", sql2struct.Special(win)))
	if c := sql2struct.Sql2structScreen(win); len(c.Objects) != 0 {
		sql2Str.Items = append(sql2Str.Items,
			widget.NewTabItemWithIcon("开始转换", theme.ViewRefreshIcon(), c),
		)
	}
	aside = widget.NewTabContainer(
		widget.NewTabItemWithIcon("首页", nil, _app.WelcomeScreen()),
		widget.NewTabItemWithIcon("Sql转Struct", nil, fyne.NewContainerWithLayout(
			layout.NewGridLayoutWithColumns(1),
			sql2Str,
		)),
		widget.NewTabItemWithIcon("Json转Struct", nil, fyne.NewContainerWithLayout(
			layout.NewGridLayoutWithColumns(1),
		)),
		widget.NewTabItemWithIcon("Map转Struct", nil, fyne.NewContainerWithLayout(
			layout.NewGridLayoutWithColumns(1),
		)),
		widget.NewTabItemWithIcon("URL编解码", nil, fyne.NewContainerWithLayout(
			layout.NewGridLayoutWithColumns(1),
		)),
		widget.NewTabItemWithIcon("日期格式化", nil, fyne.NewContainerWithLayout(
			layout.NewGridLayoutWithColumns(1),
		)),
		widget.NewTabItemWithIcon("进制转换", nil, fyne.NewContainerWithLayout(
			layout.NewGridLayoutWithColumns(1),
		)),
	)
	aside.SetTabLocation(widget.TabLocationLeading)
	return
}
