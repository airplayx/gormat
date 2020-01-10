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
	"github.com/buger/jsonparser"
	_app "gormat/app"
	"gormat/app/json2struct"
	"gormat/app/sql2struct"
	"gormat/controllers/Sql2struct"
	"net/url"
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
	IPBox := widget.NewTabContainer(ipList...)
	addBox := widget.NewScrollContainer(sql2struct.DataBase(win, options, 0))
	addBox.Hide()
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(_app.Store, func() {
			IPBox.Show()
			addBox.Hide()
		}),
		widget.NewToolbarAction(_app.SQL, func() {

		}),
		widget.NewToolbarAction(_app.JSON, func() {
			w := fyne.CurrentApp().NewWindow("Json语句转Struct")
			w.SetContent(fyne.NewContainerWithLayout(
				layout.NewGridLayout(1),
				widget.NewScrollContainer(json2struct.Screen()),
			))
			scale, _ := jsonparser.GetFloat(_app.Config, "const", "scale")
			w.Canvas().SetScale(float32(scale))
			w.Resize(fyne.Size{Width: 1000, Height: 500})
			w.CenterOnScreen()
			w.Show()
		}),
		widget.NewToolbarAction(_app.URL, func() {

		}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(_app.Option, func() {
			w := fyne.CurrentApp().NewWindow("选项")
			setting := widget.NewTabContainer(
				widget.NewTabItem("基本", sql2struct.Option(w, options)),
				widget.NewTabItem("映射", sql2struct.Reflect(w, options)),
				widget.NewTabItem("特殊转型", sql2struct.Special(w, options)),
			)
			setting.SetTabLocation(widget.TabLocationLeading)
			w.SetContent(setting)
			scale, _ := jsonparser.GetFloat(_app.Config, "const", "scale")
			w.Canvas().SetScale(float32(scale))
			w.Resize(fyne.Size{Width: 650, Height: 300})
			w.CenterOnScreen()
			w.Show()
		}),
		widget.NewToolbarAction(_app.Insert, func() {
			IPBox.Hide()
			addBox.Show()
		}),
		widget.NewToolbarAction(_app.Edit, func() {

		}),
		widget.NewToolbarAction(_app.GroupDelete, func() {
			content := widget.NewEntry()
			content.SetPlaceHolder(fmt.Sprintf("请输入 %s 确认删除当前组记录", sql2struct.CurLink[2]))
			content.OnChanged = func(text string) {
				if text == sql2struct.CurLink[2] {
					dialog.ShowInformation("操作", "删除成功", win)
				}
			}
			dialog.ShowCustom("操作", "取消", content, win)
		}),
		widget.NewToolbarAction(_app.Delete, func() {
			cnf := dialog.NewConfirm("操作", fmt.Sprintf("确定删除当前 %s 库连接记录?", sql2struct.CurLink[4]), func(b bool) {
				fmt.Println(b)
			}, win)
			cnf.SetDismissText("否")
			cnf.SetConfirmText("是")
			cnf.Show()
		}),
		widget.NewToolbarSpacer(),
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
	return fyne.NewContainerWithLayout(layout.NewAdaptiveGridLayout(1), s2sBox)
}
