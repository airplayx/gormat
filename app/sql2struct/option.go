/*
@Time : 2019/12/20 16:40
@Software: GoLand
@File : option
@Author : Bingo <airplayx@gmail.com>
*/
package sql2struct

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
	"github.com/chenhg5/collection"
	"strings"

	"gormat/common"
)

func Option(options *common.SQL2Struct) fyne.Widget {
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
	if options.JSONOmitempty {
		jsonType.SetSelected("仅生成")
	} else {
		jsonType.SetSelected("生成并包含 omitempty")
	}

	json := widget.NewCheck("json", func(on bool) {
		if !on {
			jsonType.Hide()
		} else {
			jsonType.Show()
		}
	})
	jsonType.Hidden = !collection.Collect(options.Tags).Contains("json")
	json.SetChecked(!jsonType.Hidden)

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

		},
		OnSubmit: func() {

		},
		Items: []*widget.FormItem{
			{Text: "输出文件夹", Widget: targetDir},
			{Text: "自动存为文件", Widget: autoSave},
			{Text: "标签选择", Widget: gorm},
			{Text: "", Widget: beegoOrm},
			{Text: "", Widget: xorm},
			{Text: "", Widget: json},
			{Text: "", Widget: jsonType},
			{Text: "排除表", Widget: excludeTables},
			{Text: "始终执行", Widget: tryComplete},
		},
	}
}
