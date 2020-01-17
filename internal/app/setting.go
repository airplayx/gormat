/*
@Time : 2019/12/26 11:02
@Software: GoLand
@File : setting
@Author : Bingo <airplayx@gmail.com>
*/
package app

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/buger/jsonparser"
	"gormat/configs"
)

func SettingScreen(app fyne.App, win fyne.Window) fyne.CanvasObject {
	theMe := widget.NewRadio([]string{"黑色", "白色"}, func(s string) {
		switch s {
		case "黑色":
			app.Settings().SetTheme(theme.DarkTheme())

		default:
			app.Settings().SetTheme(theme.LightTheme())
		}
	})
	switch t, _ := jsonparser.GetString(configs.Json, "const", "theme"); t {
	case "light":
		theMe.SetSelected("白色")
	default:
		theMe.SetSelected("黑色")
	}
	theMe.Horizontal = true

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
	switch scale, _ := jsonparser.GetFloat(configs.Json, "const", "scale"); scale {
	case 1.0:
		dpi.SetSelected("默认")
	case 2.0:
		dpi.SetSelected("4K")
	}
	dpi.Horizontal = true
	return fyne.NewContainerWithLayout(
		layout.NewGridLayoutWithColumns(1),
		widget.NewTabContainer(
			widget.NewTabItem("基本设置", widget.NewVBox(
				widget.NewGroup("主题",
					fyne.NewContainerWithLayout(
						layout.NewHBoxLayout(),
						theMe,
					),
				),
				widget.NewGroup("屏幕适配",
					fyne.NewContainerWithLayout(
						layout.NewHBoxLayout(),
						dpi,
					),
				),
			)),
		),
	)
}
