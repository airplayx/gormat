package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	_app "gormat/app"
	"gormat/app/tab"
	"os"
)

func main() {
	_ = os.Setenv("FYNE_FONT", `./source/font.ttf`)
	_ = os.Setenv("FYNE_SCALE", `2`)
	_ = os.Setenv("FYNE_THEME", "light")
	main := app.NewWithID("Gopher")
	icon, _ := fyne.LoadResourceFromPath("./Icon.png")
	main.SetIcon(icon)
	window := main.NewWindow("Gopher工具箱")
	window.CenterOnScreen()
	window.Resize(fyne.Size{Width: 800, Height: 450})

	result := widget.NewMultiLineEntry()
	tabs := widget.NewTabContainer(
		widget.NewTabItemWithIcon("首页", nil, _app.WelcomeScreen()),
		widget.NewTabItemWithIcon("SQL转Struct", nil, fyne.NewContainerWithLayout(
			layout.NewGridLayoutWithColumns(1),
			widget.NewTabContainer(
				widget.NewTabItem("选项", tab.Option()),
				widget.NewTabItem("数据库", tab.DataBase()),
				widget.NewTabItem("映射", tab.Reflect(window)),
				widget.NewTabItem("特殊转型", tab.Special(window)),
				widget.NewTabItemWithIcon("开始转换",
					theme.ViewRefreshIcon(),
					_app.Sql2structScreen(window, result)),
			),
		)),
		widget.NewTabItemWithIcon("JSON转Struct", nil, fyne.NewContainerWithLayout(
			layout.NewGridLayoutWithColumns(1),
		)),
	)
	tabs.SetTabLocation(widget.TabLocationLeading)

	window.SetContent(tabs)
	window.ShowAndRun()
	window.SetMaster()
}
