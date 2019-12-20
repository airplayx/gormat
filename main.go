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
	icon, _ := fyne.LoadResourceFromPath("./source/Icon.png")
	main.SetIcon(icon)
	window := main.NewWindow("Gopher工具箱")
	window.CenterOnScreen()
	window.Resize(fyne.Size{Width: 600, Height: 450})

	tabs := widget.NewTabContainer(
		widget.NewTabItemWithIcon("首页", theme.HomeIcon(), _app.WelcomeScreen()),
		widget.NewTabItemWithIcon("设置", theme.SettingsIcon(),
			fyne.NewContainerWithLayout(
				layout.NewBorderLayout(nil, nil, nil, nil),
				widget.NewTabContainer(
					widget.NewTabItem("选项", tab.Option()),
					widget.NewTabItem("数据库", tab.DataBase()),
					widget.NewTabItem("映射", tab.Reflect()),
					widget.NewTabItem("特殊转型", tab.Special()),
				),
			)),
	)
	tabs.SetTabLocation(widget.TabLocationLeading)

	window.SetContent(tabs)
	window.ShowAndRun()
	window.SetMaster()
	//log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
	//if err := common.LoadConfig(); err != nil {
	//	log.Fatalln(err.Error())
	//}
	//if err := common.NewGenTool().Gen(); err != nil {
	//	log.Fatalln(err.Error())
	//}
}
