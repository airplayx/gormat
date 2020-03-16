package app

import (
	"encoding/json"
	"errors"
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/buger/jsonparser"
	"gormat/configs"
	"gormat/internal/app/json2struct"
	"gormat/internal/app/sql2struct"
	"gormat/internal/pkg/icon"
	"gormat/pkg/Sql2struct"
	"log"
	"net/url"
	"os/exec"
	"runtime"
)

func ToolBar(win fyne.Window, ipBox *widget.TabContainer, options *Sql2struct.SQL2Struct) *widget.Toolbar {
	return widget.NewToolbar(
		widget.NewToolbarAction(icon.Insert, func() {
			w := fyne.CurrentApp().NewWindow("添加连接")
			w.SetContent(widget.NewScrollContainer(
				sql2struct.DataBase(win, w, ipBox, options, nil)),
			)
			w.Resize(fyne.Size{Width: 650, Height: 300})
			w.CenterOnScreen()
			w.Show()
		}),
		widget.NewToolbarAction(icon.Option, func() {
			w := fyne.CurrentApp().NewWindow("转换规则设置")
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
		widget.NewToolbarAction(icon.SQL, func() {
			w := fyne.CurrentApp().NewWindow("Sql语句转Struct")
			w.SetContent(fyne.NewContainerWithLayout(
				layout.NewGridLayout(1),
				widget.NewScrollContainer(sql2struct.QuickScreen()),
			))
			w.Resize(fyne.Size{Width: 1000, Height: 500})
			w.CenterOnScreen()
			w.Show()
		}),
		widget.NewToolbarAction(icon.JSON, func() {
			w := fyne.CurrentApp().NewWindow("Json语句转Struct")
			w.SetContent(fyne.NewContainerWithLayout(
				layout.NewGridLayout(1),
				widget.NewScrollContainer(json2struct.Screen()),
			))
			w.Resize(fyne.Size{Width: 1000, Height: 500})
			w.CenterOnScreen()
			w.Show()
		}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(icon.Bash, func() {
			switch runtime.GOOS {
			case "darwin":
				cmd := exec.Command(`osascript`, "-s", "h", "-e", `tell application "Terminal" to do script "echo test"`)
				if err := cmd.Start(); err != nil {
					log.Fatalln(err)
				}
			case "windows":
				cmd := exec.Command("cmd", "/C", "start", "cmd.exe")
				if err := cmd.Start(); err != nil {
					log.Fatalln(err)
				}
			case "linux":
				cmd := exec.Command("/bin/bash", "-c", `df -lh`)
				if err := cmd.Start(); err != nil {
					log.Fatalln(err)
				}
			}
			//w := fyne.CurrentApp().NewWindow("命令行相关工具")
			//w.SetContent(fyne.NewContainerWithLayout(
			//	layout.NewGridLayout(1),
			//	widget.NewScrollContainer(fyne.NewContainerWithLayout(
			//		layout.NewGridLayout(1),
			//	)),
			//))
			//w.Resize(fyne.Size{Width: 1000, Height: 500})
			//w.CenterOnScreen()
			//w.Show()
		}),
		widget.NewToolbarAction(icon.Edit, func() {
			if ipBox.Items == nil {
				return
			}
			dbBox := ipBox.CurrentTab().Content.(*widget.TabContainer)
			w := fyne.CurrentApp().NewWindow("编辑连接")
			sourceMap := options.SourceMap
			var dbIndex []int
			for k, v := range sourceMap {
				if v.Host+":"+v.Port == ipBox.CurrentTab().Text {
					for key, db := range v.Db {
						if db == dbBox.CurrentTab().Text {
							dbIndex = []int{k, key}
							goto loop
						}
					}
				}
			}
		loop:
			if dbIndex == nil {
				w.SetContent(widget.NewLabelWithStyle(`错误的连接参数`, fyne.TextAlignCenter, fyne.TextStyle{Bold: true}))
			} else {
				w.SetContent(widget.NewScrollContainer(
					sql2struct.DataBase(win, w, ipBox, options, dbIndex)))
			}
			w.Resize(fyne.Size{Width: 650, Height: 300})
			w.CenterOnScreen()
			w.Show()
		}),
		widget.NewToolbarAction(icon.GroupDelete, func() {
			if ipBox.Items == nil {
				return
			}
			content := widget.NewEntry()
			content.SetPlaceHolder("请输入 " + ipBox.CurrentTab().Text + " 确认删除当前组记录")
			content.OnChanged = func(text string) {
				if text == ipBox.CurrentTab().Text {
					sourceMap := options.SourceMap
					for k, v := range sourceMap {
						if v.Host+":"+v.Port == ipBox.CurrentTab().Text {
							options.SourceMap = append(sourceMap[:k], sourceMap[k+1:]...)
							break
						}
					}
					jsons, _ := json.Marshal(options)
					if data, err := jsonparser.Set(configs.Json, jsons, "sql2struct"); err == nil {
						configs.Json = data
						dialog.ShowInformation("操作", "保存成功", win)
						ipBox.RemoveIndex(ipBox.CurrentTabIndex())
						if ipBox.CurrentTabIndex()-1 < 0 {
							if len(ipBox.Items) == 0 {
								ipBox.Hide()
							}
							return
						}
						ipBox.SelectTabIndex(ipBox.CurrentTabIndex() - 1)
					} else {
						dialog.ShowError(errors.New(err.Error()), win)
					}
				}
			}
			dialog.ShowCustom("操作", "取消", content, win)
		}),
		widget.NewToolbarAction(icon.Delete, func() {
			if ipBox.Items == nil {
				return
			}
			dbBox := ipBox.CurrentTab().Content.(*widget.TabContainer)
			cnf := dialog.NewConfirm("操作", "确定删除当前 "+dbBox.CurrentTab().Text+" 库连接记录?", func(isDelete bool) {
				if isDelete {
					sourceMap := options.SourceMap
					for k, v := range sourceMap {
						if v.Host+":"+v.Port == ipBox.CurrentTab().Text {
							for key, db := range v.Db {
								if db == dbBox.CurrentTab().Text {
									sourceMap[k].Db = append(v.Db[:key], v.Db[key+1:]...)
									if len(sourceMap[k].Db) == 0 {
										options.SourceMap = append(sourceMap[:k], sourceMap[k+1:]...)
									}

									if dbBox.CurrentTabIndex() >= 0 && dbBox.CurrentTabIndex() < len(dbBox.Items)-1 {
										dbBox.RemoveIndex(dbBox.CurrentTabIndex())
										goto loop
									} else if dbBox.CurrentTabIndex() > 0 && dbBox.CurrentTabIndex() == len(dbBox.Items)-1 {
										dbBox.SelectTabIndex(dbBox.CurrentTabIndex() - 1)
										dbBox.RemoveIndex(dbBox.CurrentTabIndex() + 1)
										goto loop
									} else if dbBox.CurrentTabIndex() == 0 && len(dbBox.Items) == 1 {
										ipBox.RemoveIndex(ipBox.CurrentTabIndex())
										if ipBox.CurrentTabIndex()-1 < 0 {
											if len(ipBox.Items) == 0 {
												ipBox.Hide()
											}
										}
										ipBox.SelectTabIndex(ipBox.CurrentTabIndex() - 1)
										goto loop
									}

								}
							}
						}
					}
				loop:
					jsons, _ := json.Marshal(options)
					if data, err := jsonparser.Set(configs.Json, jsons, "sql2struct"); err == nil {
						configs.Json = data
						//dialog.ShowInformation("操作", "保存成功", win)
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
		widget.NewToolbarAction(icon.Github, func() {
			airPlayX, _ := url.Parse("https://github.com/airplayx/gormat")
			_ = fyne.CurrentApp().OpenURL(airPlayX)
		}),
	)
}
