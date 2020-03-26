/*
@Time : 2019/12/20 16:40
@Software: GoLand
@File : option
@Author : Bingo <airplayx@gmail.com>
*/
package sql2struct

import (
	"encoding/json"
	"errors"
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
	"github.com/buger/jsonparser"
	"github.com/chenhg5/collection"
	"gormat/configs"
	"gormat/pkg/Sql2struct"
	"strings"
)

func Option(win fyne.Window, options *Sql2struct.SQL2Struct) fyne.Widget {
	targetDir := widget.NewEntry()
	targetDir.SetText(options.TargetDir)
	autoSave := widget.NewRadio([]string{configs.Text("yes"), configs.Text("no")}, func(s string) {

	})
	autoSave.Horizontal = true
	if options.AutoSave {
		autoSave.SetSelected(configs.Text("yes"))
	} else {
		autoSave.SetSelected(configs.Text("no"))
	}

	gorm := widget.NewCheck("gorm", func(bool) {})
	gorm.SetChecked(collection.Collect(options.Tags).Contains("gorm"))

	xorm := widget.NewCheck("xorm", func(bool) {})
	xorm.SetChecked(collection.Collect(options.Tags).Contains("xorm"))

	beegoOrm := widget.NewCheck("beegoOrm", func(bool) {})
	beegoOrm.SetChecked(collection.Collect(options.Tags).Contains("beegoOrm"))

	jsonType := widget.NewSelect([]string{configs.Text("build only"), configs.Text("build and include omitempty")}, func(s string) {

	})
	if options.JsonOmitempty {
		jsonType.SetSelected(configs.Text("build and include omitempty"))
	} else {
		jsonType.SetSelected(configs.Text("build only"))
	}

	jsonT := widget.NewCheck("json", func(on bool) {
		if !on {
			jsonType.Hide()
		} else {
			jsonType.Show()
		}
	})
	jsonType.Hidden = !collection.Collect(options.Tags).Contains("json")
	jsonT.SetChecked(!jsonType.Hidden)

	excludeTables := widget.NewMultiLineEntry()
	excludeTables.SetPlaceHolder(configs.Text("carriage return"))
	excludeTables.SetText(strings.Join(options.ExcludeTables, "\n"))
	tryComplete := widget.NewRadio([]string{configs.Text("yes"), configs.Text("no")}, func(s string) {

	})
	tryComplete.Horizontal = true
	if options.TryComplete {
		tryComplete.SetSelected(configs.Text("yes"))
	} else {
		tryComplete.SetSelected(configs.Text("no"))
	}
	return &widget.Form{
		OnCancel: func() {
			win.Close()
		},
		OnSubmit: func() {
			options.TargetDir = targetDir.Text
			options.AutoSave = autoSave.Selected == configs.Text("yes")
			options.Tags = []string{}
			if gorm.Checked {
				options.Tags = append(options.Tags, "gorm")
			}
			if xorm.Checked {
				options.Tags = append(options.Tags, "xorm")
			}
			if beegoOrm.Checked {
				options.Tags = append(options.Tags, "beegoOrm")
			}
			if jsonT.Checked {
				options.Tags = append(options.Tags, "json")
			}
			options.JsonOmitempty = jsonType.Selected == configs.Text("build and include omitempty")
			options.ExcludeTables = strings.Split(excludeTables.Text, "\n")
			options.TryComplete = tryComplete.Selected == configs.Text("yes")

			jsons, _ := json.Marshal(options)
			if data, err := jsonparser.Set(configs.Json, jsons, "sql2struct"); err == nil {
				configs.Json = data
				dialog.ShowInformation(configs.Text("info"), configs.Text("save ok"), win)
			} else {
				dialog.ShowError(errors.New(err.Error()), win)
			}
		},
		Items: []*widget.FormItem{
			{Text: configs.Text("auto save files"), Widget: autoSave},
			{Text: configs.Text("output folder"), Widget: targetDir},
			{Text: configs.Text("tags"), Widget: gorm},
			//{Text: "", Widget: beegoOrm},
			{Text: "", Widget: xorm},
			{Text: "", Widget: jsonT},
			{Text: "", Widget: jsonType},
			{Text: configs.Text("exclusion table"), Widget: excludeTables},
			{Text: configs.Text("always execute"), Widget: tryComplete},
		},
	}
}
