package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"github.com/buger/jsonparser"
	"gormat/configs"
	_app "gormat/internal/app"
	"gormat/internal/pkg/icon"
	"io/ioutil"
	"os"
)

func main() {
	jsonparser.EachKey(configs.Json,
		func(i int, bytes []byte, valueType jsonparser.ValueType, e error) {
			font, _ := jsonparser.GetString(bytes, "font")
			if font != "" {
				_ = os.Setenv("FYNE_FONT", font)
			}
			theme, _ := jsonparser.GetString(bytes, "theme")
			_ = os.Setenv("FYNE_THEME", theme)
			scale, _ := jsonparser.GetString(bytes, "scale")
			_ = os.Setenv("FYNE_SCALE", scale)
		}, []string{"const"})
	main := app.NewWithID("Gopher")
	main.SetIcon(icon.Ico)
	window := main.NewWindow("Gormat")
	window.CenterOnScreen()
	window.Resize(fyne.Size{Width: 1200, Height: 720})
	window.SetContent(_app.Container(main, window))
	window.SetOnClosed(func() {
		_ = ioutil.WriteFile(configs.CustomFile, configs.Json, os.ModePerm)
	})
	window.SetMaster()
	window.ShowAndRun()
}

/**
vendor/fyne.io/fyne/widget/tabcontainer.go:547
vendor/fyne.io/fyne/widget/tabcontainer.go:193
vendor/fyne.io/fyne/widget/tabcontainer.go:411
vendor/fyne.io/fyne/widget/form.go:94
vendor/fyne.io/fyne/dialog/confirm.go:28
*/
