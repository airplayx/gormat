/*
@Time : 2019/12/26 11:02
@Software: GoLand
@File : setting
@Author : Bingo <airplayx@gmail.com>
*/
package app

import (
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/buger/jsonparser"
	"gormat/configs"
	"os"
)

func SettingScreen(app fyne.App, win fyne.Window) fyne.CanvasObject {
	theMe := widget.NewRadio([]string{configs.Text("black"), configs.Text("light")}, func(s string) {
		switch s {
		case configs.Text("black"):
			app.Settings().SetTheme(theme.DarkTheme())

		default:
			app.Settings().SetTheme(theme.LightTheme())
		}
	})

	switch t, _ := jsonparser.GetString(configs.Json, "const", "theme"); t {
	case configs.Text("light"):
		theMe.SetSelected(configs.Text("light"))
	default:
		theMe.SetSelected(configs.Text("black"))
	}
	theMe.Horizontal = true

	dpi := widget.NewRadio([]string{configs.Text("default"), "4K"}, func(s string) {
		scale := "1.0"
		switch s {
		case "4K":
			scale = "2.0"
		}
		_ = os.Setenv("FYNE_SCALE", scale)
		if data, err := jsonparser.Set(configs.Json, []byte("\""+scale+"\""), "const", "scale"); err == nil {
			configs.Json = data
		}
	})
	switch scale, _ := jsonparser.GetString(configs.Json, "const", "scale"); scale {
	case "1.0":
		dpi.SetSelected(configs.Text("default"))
	case "2.0":
		dpi.SetSelected("4K")
	}
	dpi.Horizontal = true

	language := widget.NewRadio([]string{"中文", "English"}, func(s string) {
		//defer app.Settings().SetTheme(theme.LightTheme())
		lang := "en"
		switch s {
		case "中文":
			lang = "zh"
		}
		if data, err := jsonparser.Set(configs.Json, []byte("\""+lang+"\""), "const", "language"); err == nil {
			configs.Json = data
		}
	})
	l, _ := jsonparser.GetString(configs.Json, "const", "language")
	if l == "zh" {
		language.SetSelected("中文")
	} else {
		language.SetSelected("English")
	}
	language.Horizontal = true

	return &widget.Form{
		OnCancel: func() {
			win.Close()
		},
		OnSubmit: func() {
			dialog.ShowInformation(configs.Text("info"), configs.Text("effective after restart"), win)
		},
		Items: []*widget.FormItem{
			{Text: configs.Text("screen"), Widget: fyne.NewContainerWithLayout(
				layout.NewHBoxLayout(),
				dpi,
			)},
			{Text: configs.Text("language"), Widget: fyne.NewContainerWithLayout(
				layout.NewHBoxLayout(),
				language,
			)},
			//{Text: configs.Text("theme"), Widget: fyne.NewContainerWithLayout(
			//	layout.NewHBoxLayout(),
			//	theMe,
			//)},
		},
	}
}
