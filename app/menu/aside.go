/*
@Time : 2019/12/23 10:24
@Software: GoLand
@File : aside
@Author : Bingo <airplayx@gmail.com>
*/
package menu

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	_app "gormat/app"
	"gormat/app/json2struct"
	"gormat/app/sql2struct"
	"gormat/controllers/Sql2struct"
	"net/url"
)

func Aside(app fyne.App, win fyne.Window) (aside *widget.TabContainer) {
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
	IPBox := widget.NewTabContainer(ipList...)
	addBox := widget.NewScrollContainer(
		sql2struct.DataBase(win, options, 0),
	)
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(_app.Store, func() {
			IPBox.Show()
			addBox.Hide()
		}),
		widget.NewToolbarAction(_app.Insert, func() {
			IPBox.Hide()
			addBox.Show()
		}),
		widget.NewToolbarAction(_app.SQL, func() {
			w := fyne.CurrentApp().NewWindow("SQL语句转Struct")
			w.SetContent(fyne.NewContainerWithLayout(
				layout.NewGridLayout(1),
				widget.NewScrollContainer(json2struct.Screen()),
			))
			w.Resize(fyne.Size{Width: 1000, Height: 500})
			w.CenterOnScreen()
			w.Show()
		}),
		//widget.NewToolbarSeparator(),
		widget.NewToolbarAction(_app.GroupDelete, func() {
			content := widget.NewEntry()
			content.SetPlaceHolder("输入 yes 删除当前组")
			content.OnChanged = func(text string) {
				if text == "yes" {
					dialog.ShowInformation("操作", "删除成功", win)
				}
			}
			dialog.ShowCustom("操作", "取消", content, win)
		}),
		widget.NewToolbarAction(_app.Delete, func() {
			cnf := dialog.NewConfirm("操作", "确定删除当前记录?", func(b bool) {
				fmt.Println(b)
			}, win)
			cnf.SetDismissText("否")
			cnf.SetConfirmText("是")
			cnf.Show()
		}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(_app.Setting, func() {
			w := fyne.CurrentApp().NewWindow("设置")
			setting := widget.NewTabContainer(
				widget.NewTabItem("基本", sql2struct.Option(w, options)),
				widget.NewTabItem("映射", sql2struct.Reflect(w, options)),
				widget.NewTabItem("特殊转型", sql2struct.Special(w, options)),
			)
			setting.SetTabLocation(widget.TabLocationLeading)
			w.SetContent(setting)
			w.Resize(fyne.Size{Width: 650, Height: 300})
			w.CenterOnScreen()
			w.Show()
		}),
		widget.NewToolbarAction(_app.Info, func() {
			airPlayX, _ := url.Parse("http://airplayx.com/gopher-tool")
			_ = fyne.CurrentApp().OpenURL(airPlayX)
		}),
	)
	s2sBox := fyne.NewContainerWithLayout(
		layout.NewBorderLayout(toolbar, nil, nil, nil),
		toolbar,
		addBox,
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
