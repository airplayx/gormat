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
	"gormat/app/config"
	"gormat/internal/Sql2struct"
	"strings"
)

func Option(win fyne.Window, options *Sql2struct.SQL2Struct) fyne.Widget {
	targetDir := widget.NewEntry()
	targetDir.SetText(options.TargetDir)
	autoSave := widget.NewRadio([]string{"是", "否"}, func(s string) {

	})
	autoSave.Horizontal = true
	if options.AutoSave {
		autoSave.SetSelected("是")
	} else {
		autoSave.SetSelected("否")
	}

	gorm := widget.NewCheck("gorm", func(bool) {})
	gorm.SetChecked(collection.Collect(options.Tags).Contains("gorm"))

	xorm := widget.NewCheck("xorm", func(bool) {})
	xorm.SetChecked(collection.Collect(options.Tags).Contains("xorm"))

	beegoOrm := widget.NewCheck("beegoOrm", func(bool) {})
	beegoOrm.SetChecked(collection.Collect(options.Tags).Contains("beegoOrm"))

	jsonType := widget.NewSelect([]string{"仅生成", "生成并包含 omitempty"}, func(s string) {

	})
	if options.JsonOmitempty {
		jsonType.SetSelected("生成并包含 omitempty")
	} else {
		jsonType.SetSelected("仅生成")
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
	excludeTables.SetPlaceHolder("多个数据表以回车换行")
	excludeTables.SetText(strings.Join(options.ExcludeTables, "\n"))
	tryComplete := widget.NewRadio([]string{"是", "否"}, func(s string) {

	})
	tryComplete.Horizontal = true
	if options.TryComplete {
		tryComplete.SetSelected("是")
	} else {
		tryComplete.SetSelected("否")
	}
	return &widget.Form{
		OnCancel: func() {
			win.Close()
		},
		OnSubmit: func() {
			options.TargetDir = targetDir.Text
			options.AutoSave = autoSave.Selected == "是"
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
			options.JsonOmitempty = jsonType.Selected == "生成并包含 omitempty"
			options.ExcludeTables = strings.Split(excludeTables.Text, "\n")
			options.TryComplete = tryComplete.Selected == "是"

			jsons, _ := json.Marshal(options)
			if data, err := jsonparser.Set(config.Setting, jsons, "sql2struct"); err == nil {
				config.Setting = data
				dialog.ShowInformation("成功", "保存成功", win)
			} else {
				dialog.ShowError(errors.New(err.Error()), win)
			}
		},
		Items: []*widget.FormItem{
			{Text: "输出文件夹", Widget: targetDir},
			{Text: "自动存文件", Widget: autoSave},
			{Text: "标签选择", Widget: gorm},
			//{Text: "", Widget: beegoOrm},
			{Text: "", Widget: xorm},
			{Text: "", Widget: jsonT},
			{Text: "", Widget: jsonType},
			{Text: "排除表", Widget: excludeTables},
			{Text: "始终执行", Widget: tryComplete},
		},
	}
}
