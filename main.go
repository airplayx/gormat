package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"github.com/buger/jsonparser"
	"gormat/app"
	"gormat/app/menu"
	"io/ioutil"
	"os"
)

func main() {
	defer func() {
		_ = ioutil.WriteFile(_app.ConFile, _app.Config, os.ModePerm)
	}()
	jsonparser.EachKey(_app.Config,
		func(i int, bytes []byte, valueType jsonparser.ValueType, e error) {
			font, _ := jsonparser.GetString(bytes, "font")
			_ = os.Setenv("FYNE_FONT", font)
			theme, _ := jsonparser.GetString(bytes, "theme")
			_ = os.Setenv("FYNE_THEME", theme)
		}, []string{"const"})
	main := app.NewWithID("Gopher")
	main.SetIcon(fyne.NewStaticResource("ico", _app.Ico))
	window := main.NewWindow("Gopher工具箱")
	scale, _ := jsonparser.GetFloat(_app.Config, "const", "scale")
	window.Canvas().SetScale(float32(scale))
	window.CenterOnScreen()
	window.Resize(fyne.Size{Width: 1000, Height: 610})
	window.SetContent(menu.Aside(main, window))
	window.SetMaster()
	window.ShowAndRun()
}
