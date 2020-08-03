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
	s2s "gormat/internal/app/sql2struct"
	"gormat/internal/pkg/icon"
	"gormat/pkg/sql2struct"
	"log"
	"net/url"
	"os/exec"
	"runtime"
)

//ToolBar the main toolbar
func ToolBar(app fyne.App, win fyne.Window, ipBox *widget.TabContainer, options *sql2struct.SQL2Struct) *widget.Toolbar {
	return widget.NewToolbar(
		widget.NewToolbarAction(icon.Insert, func() {
			var dbIndex []int
			if ipBox.Items != nil && ipBox.Visible() {
				sourceMap := options.SourceMap
				for k := range sourceMap {
					if k == ipBox.CurrentTabIndex() {
						dbIndex = []int{k}
						break
					}
				}
			}
			w := fyne.CurrentApp().NewWindow(configs.Text("add link"))
			w.SetContent(widget.NewScrollContainer(
				s2s.DataBase(win, w, ipBox, options, dbIndex)),
			)
			w.Resize(fyne.Size{Width: 600, Height: 300})
			w.CenterOnScreen()
			w.Show()
		}),
		widget.NewToolbarAction(icon.Setting, func() {
			w := fyne.CurrentApp().NewWindow(configs.Text("rules"))
			setting := widget.NewTabContainer(
				widget.NewTabItem(configs.Text("base"), s2s.Option(w, options)),
				widget.NewTabItem(configs.Text("mapping"), s2s.Reflect(w, options)),
				widget.NewTabItem(configs.Text("special"), s2s.Special(w, options)),
				widget.NewTabItem(configs.Text("others"), SettingScreen(app, w)),
			)
			setting.SetTabLocation(widget.TabLocationLeading)
			w.SetContent(setting)
			w.Resize(fyne.Size{Width: 500, Height: 300})
			w.CenterOnScreen()
			w.Show()
		}),
		widget.NewToolbarAction(icon.SQL, func() {
			w := fyne.CurrentApp().NewWindow(configs.Text("sql to struct"))
			w.SetContent(fyne.NewContainerWithLayout(
				layout.NewGridLayout(1),
				widget.NewScrollContainer(s2s.QuickScreen()),
			))
			w.Resize(fyne.Size{Width: 1000, Height: 500})
			w.CenterOnScreen()
			w.Show()
		}),
		widget.NewToolbarAction(icon.JSON, func() {
			w := fyne.CurrentApp().NewWindow(configs.Text("json to struct"))
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
		}),
		widget.NewToolbarAction(icon.Edit, func() {
			if ipBox.Items == nil {
				return
			}
			dbBox := ipBox.CurrentTab().Content.(*widget.TabContainer)
			win.Canvas().Refresh(dbBox)
			w := fyne.CurrentApp().NewWindow(configs.Text("edit link"))
			sourceMap := options.SourceMap
			var dbIndex []int
			for k, v := range sourceMap {
				if k == ipBox.CurrentTabIndex() {
					for key := range v.Db {
						if key == dbBox.CurrentTabIndex() {
							dbIndex = []int{k, key}
							goto loop
						}
					}
				}
			}
		loop:
			if dbIndex == nil {
				w.SetContent(widget.NewLabelWithStyle(configs.Text("bad link parameters"), fyne.TextAlignCenter, fyne.TextStyle{Bold: true}))
			} else {
				w.SetContent(widget.NewScrollContainer(
					s2s.DataBase(win, w, ipBox, options, dbIndex)))
			}
			w.Resize(fyne.Size{Width: 600, Height: 300})
			w.CenterOnScreen()
			w.Show()
		}),
		widget.NewToolbarAction(icon.GroupDelete, func() {
			if ipBox.Items == nil || ipBox.CurrentTab() == nil {
				return
			}
			content := widget.NewEntry()
			content.SetPlaceHolder(configs.Text("Please input %s to delete the current group record", ipBox.CurrentTab().Text))
			content.OnChanged = func(text string) {
				defer ipBox.SelectTabIndex(ipBox.CurrentTabIndex())
				if text == ipBox.CurrentTab().Text {
					sourceMap := options.SourceMap
					for k := range sourceMap {
						if k == ipBox.CurrentTabIndex() {
							options.SourceMap = append(sourceMap[:k], sourceMap[k+1:]...)
							break
						}
					}
					jsons, _ := json.Marshal(options)
					if data, err := jsonparser.Set(configs.JSON, jsons, "sql2struct"); err == nil {
						configs.JSON = data
						dialog.ShowInformation(configs.Text("action"), configs.Text("save ok"), win)
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
			dialog.ShowCustom(configs.Text("action"), configs.Text("cancel"), content, win)
		}),
		widget.NewToolbarAction(icon.Delete, func() {
			if ipBox.Items == nil || ipBox.CurrentTab() == nil {
				return
			}
			dbBox := ipBox.CurrentTab().Content.(*widget.TabContainer)
			if dbBox.Items == nil || dbBox.CurrentTab() == nil {
				return
			}
			cnf := dialog.NewConfirm(configs.Text("action"), configs.Text("Please confirm to delete the current: %s ?", dbBox.CurrentTab().Text), func(isDelete bool) {
				if isDelete {
					sourceMap := options.SourceMap
					for k, v := range sourceMap {
						if k == ipBox.CurrentTabIndex() {
							for key := range v.Db {
								if key == dbBox.CurrentTabIndex() {
									sourceMap[k].Db = append(v.Db[:key], v.Db[key+1:]...)
									if len(sourceMap[k].Db) == 0 {
										options.SourceMap = append(sourceMap[:k], sourceMap[k+1:]...)
									}
									if dbBox.CurrentTabIndex() >= 0 && dbBox.CurrentTabIndex() < len(dbBox.Items)-1 { //delete the 0~n db
										dbBox.RemoveIndex(dbBox.CurrentTabIndex())
										dbBox.SelectTabIndex(dbBox.CurrentTabIndex())
										goto loop
									} else if dbBox.CurrentTabIndex() > 0 && dbBox.CurrentTabIndex() == len(dbBox.Items)-1 { //delete the last db
										dbBox.SelectTabIndex(dbBox.CurrentTabIndex() - 1)
										dbBox.RemoveIndex(dbBox.CurrentTabIndex() + 1)
										goto loop
									} else if dbBox.CurrentTabIndex() == 0 && len(dbBox.Items) == 1 { //delete only one db
										ipBox.RemoveIndex(ipBox.CurrentTabIndex())
										if len(ipBox.Items) == 0 {
											ipBox.Hide()
										} else {
											if ipBox.CurrentTabIndex() >= len(ipBox.Items) {
												ipBox.SelectTabIndex(ipBox.CurrentTabIndex() - 1)
											} else {
												ipBox.SelectTabIndex(ipBox.CurrentTabIndex())
											}
										}
										goto loop
									}
								}
							}
						}
					}
				loop:
					defer win.Canvas().Refresh(ipBox)
					jsons, _ := json.Marshal(options)
					if data, err := jsonparser.Set(configs.JSON, jsons, "sql2struct"); err == nil {
						configs.JSON = data
					} else {
						dialog.ShowError(errors.New(err.Error()), win)
					}
				}
			}, win)
			cnf.SetDismissText(configs.Text("no"))
			cnf.SetConfirmText(configs.Text("yes"))
			cnf.Show()
		}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(icon.Github, func() {
			airPlayX, _ := url.Parse("https://github.com/airplayx/gormat")
			_ = fyne.CurrentApp().OpenURL(airPlayX)
		}),
	)
}
