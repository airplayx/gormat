//go:generate statik -src=../assets/fonts
//go:generate go fmt statik/statik.go
package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"github.com/buger/jsonparser"
	"github.com/rakyll/statik/fs"
	_ "gormat/cmd/statik"
	"gormat/configs"
	_app "gormat/internal/app"
	"gormat/internal/pkg/icon"
	"io/ioutil"
	"log"
	"os"
)

var tmpFile, _ = ioutil.TempFile("", "*")

func init() {
	fileSystem, err := fs.New()
	if err != nil {
		log.Fatalln(err.Error())
	}
	file, err := fileSystem.Open("/miniHei.ttf")
	if err != nil {
		log.Fatalln(err.Error())
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, _ = tmpFile.Write(content)
}

func main() {
	defer func() {
		_ = tmpFile.Close()
		_ = os.Remove(tmpFile.Name())
	}()
	jsonparser.EachKey(configs.JSON,
		func(i int, bytes []byte, valueType jsonparser.ValueType, e error) {
			_ = os.Setenv("FYNE_FONT", tmpFile.Name())
			theme, _ := jsonparser.GetString(bytes, "theme")
			_ = os.Setenv("FYNE_THEME", theme)
			scale, _ := jsonparser.GetString(bytes, "scale")
			_ = os.Setenv("FYNE_SCALE", scale)
		}, []string{"const"})
	main := app.NewWithID("Gopher")
	main.SetIcon(icon.Ico)
	window := main.NewWindow("Gormat")
	window.CenterOnScreen()
	window.Resize(fyne.Size{Width: 1300, Height: 700})
	window.SetContent(_app.Container(main, window))
	window.SetOnClosed(func() {
		_ = ioutil.WriteFile(configs.CustomFile, configs.JSON, os.ModePerm)
	})
	window.SetMaster()
	window.ShowAndRun()
}

/**
vendor/fyne.io/fyne/widget/tabcontainer.go:568
vendor/fyne.io/fyne/widget/tabcontainer.go:380
vendor/fyne.io/fyne/widget/tabcontainer.go:430
vendor/fyne.io/fyne/widget/form.go:121
vendor/fyne.io/fyne/widget/form.go:125
vendor/fyne.io/fyne/dialog/confirm.go:28
*/
