package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"github.com/buger/jsonparser"
	"gormat/app"
	"gormat/app/config"
	"io/ioutil"
	"os"
)

func main() {
	jsonparser.EachKey(config.Setting,
		func(i int, bytes []byte, valueType jsonparser.ValueType, e error) {
			font, _ := jsonparser.GetString(bytes, "font")
			_ = os.Setenv("FYNE_FONT", font)
			theme, _ := jsonparser.GetString(bytes, "theme")
			_ = os.Setenv("FYNE_THEME", theme)
		}, []string{"const"})
	main := app.NewWithID("Gopher")
	main.SetIcon(config.Ico)
	window := main.NewWindow("Gopher工具箱")
	scale, _ := jsonparser.GetFloat(config.Setting, "const", "scale")
	window.Canvas().SetScale(float32(scale))
	window.CenterOnScreen()
	window.Resize(fyne.Size{Width: 1366, Height: 650})
	window.SetContent(_app.Container(main, window))
	window.SetOnClosed(func() {
		_ = ioutil.WriteFile(config.File, config.Setting, os.ModePerm)
	})
	window.SetMaster()
	window.ShowAndRun()
}
