package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"gormat/app/menu"
	"os"
)

func main() {
	_ = os.Setenv("FYNE_FONT", `./source/font.ttf`)
	_ = os.Setenv("FYNE_SCALE", `1`)
	_ = os.Setenv("FYNE_THEME", "light")
	main := app.NewWithID("Gopher")
	icon, _ := fyne.LoadResourceFromPath("./Icon.png")
	main.SetIcon(icon)
	window := main.NewWindow("Gopher工具箱")
	window.CenterOnScreen()
	window.Resize(fyne.Size{Width: 1000, Height: 450})
	window.SetContent(menu.Aside(window))
	window.ShowAndRun()
	window.SetMaster()
}
