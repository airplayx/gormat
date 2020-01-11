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
	"gormat/app/sql2struct"
	"gormat/controllers/Sql2struct"
)

func Aside(app fyne.App, win fyne.Window) *fyne.Container {
	var options = Sql2struct.Configs()
	var ipList []*widget.TabItem
	for _, v := range options.SourceMap {
		var dbList []*widget.TabItem
		for _, curDb := range v.Db {
			dbList = append(dbList, widget.NewTabItemWithIcon(
				curDb, _app.Database,
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
	iPBox := widget.NewTabContainer(ipList...)
	tooBar := sql2struct.ToolBar(win, options)
	s2sBox := fyne.NewContainerWithLayout(
		layout.NewBorderLayout(tooBar, nil, nil, nil),
		tooBar,
		_app.WelcomeScreen(),
	)
	if len(iPBox.Items) > 0 {
		iPBox.SetTabLocation(widget.TabLocationLeading)
		s2sBox.AddObject(iPBox)
	}
	return fyne.NewContainerWithLayout(layout.NewAdaptiveGridLayout(1), s2sBox)
}
