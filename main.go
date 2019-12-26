package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"github.com/buger/jsonparser"
	"gormat/app"
	"gormat/app/menu"
	"os"
)

func main() {
	jsonparser.EachKey(_app.Config,
		func(i int, bytes []byte, valueType jsonparser.ValueType, e error) {
			font, _ := jsonparser.GetString(bytes, "font")
			_ = os.Setenv("FYNE_FONT", font)
			theme, _ := jsonparser.GetString(bytes, "theme")
			_ = os.Setenv("FYNE_THEME", theme)
			scale, _ := jsonparser.GetString(bytes, "scale")
			_ = os.Setenv("FYNE_SCALE", scale)
		}, []string{"const"})
	main := app.NewWithID("Gopher")
	main.SetIcon(fyne.NewStaticResource("ico", _app.Ico))
	window := main.NewWindow("Gopher工具箱")
	window.CenterOnScreen()
	window.Resize(fyne.Size{Width: 1366, Height: 768})
	window.SetContent(menu.Aside(main, window))
	window.ShowAndRun()
	window.SetMaster()
}
