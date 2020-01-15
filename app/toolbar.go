package _app

import (
	"encoding/json"
	"errors"
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/buger/jsonparser"
	"gormat/app/config"
	"gormat/app/json2struct"
	"gormat/app/sql2struct"
	"gormat/internal/Sql2struct"
	"net/url"
)

func ToolBar(win fyne.Window, ipBox, dbBox *widget.TabContainer, options *Sql2struct.SQL2Struct) *widget.Toolbar {
	return widget.NewToolbar(
		widget.NewToolbarAction(config.Insert, func() {
			w := fyne.CurrentApp().NewWindow("添加连接")
			w.SetContent(widget.NewScrollContainer(sql2struct.DataBase(w, options, -1)))
			scale, _ := jsonparser.GetFloat(config.Setting, "const", "scale")
			w.Canvas().SetScale(float32(scale))
			w.Resize(fyne.Size{Width: 650, Height: 300})
			w.CenterOnScreen()
			w.Show()
		}),
		widget.NewToolbarAction(config.Option, func() {
			w := fyne.CurrentApp().NewWindow("转换规则设置")
			setting := widget.NewTabContainer(
				widget.NewTabItem("基本", sql2struct.Option(w, options)),
				widget.NewTabItem("映射", sql2struct.Reflect(w, options)),
				widget.NewTabItem("特殊转型", sql2struct.Special(w, options)),
			)
			setting.SetTabLocation(widget.TabLocationLeading)
			w.SetContent(setting)
			scale, _ := jsonparser.GetFloat(config.Setting, "const", "scale")
			w.Canvas().SetScale(float32(scale))
			w.Resize(fyne.Size{Width: 650, Height: 300})
			w.CenterOnScreen()
			w.Show()
		}),
		widget.NewToolbarAction(config.SQL, func() {
			w := fyne.CurrentApp().NewWindow("Sql语句转Struct")
			w.SetContent(fyne.NewContainerWithLayout(
				layout.NewGridLayout(1),
				widget.NewScrollContainer(sql2struct.QuickScreen()),
			))
			scale, _ := jsonparser.GetFloat(config.Setting, "const", "scale")
			w.Canvas().SetScale(float32(scale))
			w.Resize(fyne.Size{Width: 1000, Height: 500})
			w.CenterOnScreen()
			w.Show()
		}),
		widget.NewToolbarAction(config.JSON, func() {
			w := fyne.CurrentApp().NewWindow("Json语句转Struct")
			w.SetContent(fyne.NewContainerWithLayout(
				layout.NewGridLayout(1),
				widget.NewScrollContainer(json2struct.Screen()),
			))
			scale, _ := jsonparser.GetFloat(config.Setting, "const", "scale")
			w.Canvas().SetScale(float32(scale))
			w.Resize(fyne.Size{Width: 1000, Height: 500})
			w.CenterOnScreen()
			w.Show()
		}),
		widget.NewToolbarAction(config.URL, func() {
			w := fyne.CurrentApp().NewWindow("Url相关工具")
			w.SetContent(fyne.NewContainerWithLayout(
				layout.NewGridLayout(1),
				widget.NewScrollContainer(fyne.NewContainerWithLayout(
					layout.NewGridLayout(1),
				)),
			))
			scale, _ := jsonparser.GetFloat(config.Setting, "const", "scale")
			w.Canvas().SetScale(float32(scale))
			w.Resize(fyne.Size{Width: 1000, Height: 500})
			w.CenterOnScreen()
			w.Show()
		}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(config.Edit, func() {
			if len(sql2struct.CurLink) == 0 {
				return
			}

		}),
		widget.NewToolbarAction(config.GroupDelete, func() {
			if len(sql2struct.CurLink) == 0 {
				return
			}
			content := widget.NewEntry()
			content.SetPlaceHolder(fmt.Sprintf("请输入 %s 确认删除当前组记录", sql2struct.CurLink[2]))
			content.OnChanged = func(text string) {
				if text == sql2struct.CurLink[2] {
					sourceMap := options.SourceMap
					for k, v := range sourceMap {
						if v.Host == sql2struct.CurLink[2] {
							options.SourceMap = append(sourceMap[:k], sourceMap[k+1:]...)
							break
						}
					}
					jsons, _ := json.Marshal(options)
					if data, err := jsonparser.Set(config.Setting, jsons, "sql2struct"); err == nil {
						config.Setting = data
						dialog.ShowInformation("操作", "保存成功", win)
						ipBox.RemoveIndex(ipBox.CurrentTabIndex())
						if ipBox.CurrentTabIndex()-1 < 0 {
							if len(ipBox.Items) == 0 {
								ipBox.Hide()
							}
							return
						}
						ipBox.SelectTabIndex(ipBox.CurrentTabIndex() - 1)
						ipBox.Refresh()
					} else {
						dialog.ShowError(errors.New(err.Error()), win)
					}
				}
			}
			dialog.ShowCustom("操作", "取消", content, win)
		}),
		widget.NewToolbarAction(config.Delete, func() {
			if len(sql2struct.CurLink) == 0 {
				return
			}
			cnf := dialog.NewConfirm("操作", fmt.Sprintf("确定删除当前 %s 库连接记录?", sql2struct.CurLink[4]), func(b bool) {
				if b {
					sourceMap := options.SourceMap
					for k, v := range sourceMap {
						if v.Host == sql2struct.CurLink[2] {
							for key, db := range v.Db {
								if db == sql2struct.CurLink[4] {
									sourceMap[k].Db = append(v.Db[:key], v.Db[key+1:]...)
									if len(sourceMap[k].Db) == 0 {
										options.SourceMap = append(sourceMap[:k], sourceMap[k+1:]...)
									}
									dbBox.RemoveIndex(dbBox.CurrentTabIndex())
									if dbBox.CurrentTabIndex()-1 < 0 {
										if len(dbBox.Items) == 0 {
											dbBox.Hide()
											ipBox.RemoveIndex(ipBox.CurrentTabIndex())
											if ipBox.CurrentTabIndex()-1 < 0 {
												if len(ipBox.Items) == 0 {
													ipBox.Hide()
												}
												goto loop
											}
											ipBox.SelectTabIndex(ipBox.CurrentTabIndex() - 1)
											ipBox.Refresh()
										}
										goto loop
									}
									dbBox.SelectTabIndex(dbBox.CurrentTabIndex() - 1)
									dbBox.Refresh()
								}
							}
						}
					}
				loop:
					jsons, _ := json.Marshal(options)
					if data, err := jsonparser.Set(config.Setting, jsons, "sql2struct"); err == nil {
						config.Setting = data
						dialog.ShowInformation("操作", "保存成功", win)
					} else {
						dialog.ShowError(errors.New(err.Error()), win)
					}
				}
			}, win)
			cnf.SetDismissText("否")
			cnf.SetConfirmText("是")
			cnf.Show()
		}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(config.Info, func() {
			airPlayX, _ := url.Parse("http://airplayx.com/gopher-tool")
			_ = fyne.CurrentApp().OpenURL(airPlayX)
		}),
	)
}
