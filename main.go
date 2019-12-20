package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	_app "gormat/app"
	"gormat/common"
	"log"
)

func main() {
	a := app.NewWithID("gormat")
	a.SetIcon(theme.FyneLogo())
	a.Settings().SetTheme(theme.LightTheme())
	w := a.NewWindow("DataBase to Gorm struct")
	w.SetMaster()

	tabs := widget.NewTabContainer(
		widget.NewTabItemWithIcon("Welcome", theme.HomeIcon(), _app.WelcomeScreen(500, 350)),
		widget.NewTabItemWithIcon("Setting", theme.SettingsIcon(), fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, nil, nil, nil),
			widget.NewTabContainer(
				widget.NewTabItem("option", _app.OptionTab()),
				widget.NewTabItem("database", _app.DataBaseTab()),
				widget.NewTabItem("reflect", _app.ReflectTab()),
				widget.NewTabItem("special", _app.SpecialTab()),
			),
		)),
	)
	tabs.SetTabLocation(widget.TabLocationLeading)

	w.SetContent(tabs)
	w.ShowAndRun()
	return
	//log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
	if err := common.LoadConfig(); err != nil {
		log.Fatalln(err.Error())
	}
	if err := common.NewGenTool().Gen(); err != nil {
		log.Fatalln(err.Error())
	}
}
