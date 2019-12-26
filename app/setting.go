/*
@Time : 2019/12/26 11:02
@Software: GoLand
@File : setting
@Author : Bingo <airplayx@gmail.com>
*/
package _app

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func SettingScreen(app fyne.App, win fyne.Window) fyne.CanvasObject {
	theme := widget.NewRadio([]string{"黑色", "白色"}, func(s string) {
		switch s {
		case "黑色":
			app.Settings().SetTheme(theme.DarkTheme())

		default:
			app.Settings().SetTheme(theme.LightTheme())
		}
	})
	theme.SetSelected("白色")
	theme.Horizontal = true

	dpi := widget.NewRadio([]string{"默认" /*"全屏",*/, "4K"}, func(s string) {
		//win.SetFullScreen(false)
		switch s {
		case "4K":
			win.Canvas().SetScale(2)
		//case "全屏":
		//	win.SetFullScreen(true)
		default:
			win.Canvas().SetScale(1)
		}
	})
	dpi.SetSelected("默认")
	dpi.Horizontal = true
	return fyne.NewContainerWithLayout(
		layout.NewGridLayoutWithColumns(1),
		widget.NewTabContainer(
			widget.NewTabItem("基本设置", widget.NewVBox(
				widget.NewGroup("主题",
					fyne.NewContainerWithLayout(
						layout.NewHBoxLayout(),
						theme,
					),
				),
				widget.NewGroup("DPI适应",
					fyne.NewContainerWithLayout(
						layout.NewHBoxLayout(),
						dpi,
					),
				),
			)),
		),
	)
}
