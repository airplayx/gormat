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
)

func Option() fyne.Widget {
	targetDir := widget.NewEntry()
	targetDir.SetText("./models")
	autoSave := widget.NewRadio([]string{"是", "否"}, func(s string) {

	})
	autoSave.Horizontal = true
	autoSave.SetSelected("是")

	gorm := widget.NewCheck("gorm", func(bool) {})
	gorm.SetChecked(true)

	xorm := widget.NewCheck("xorm", func(bool) {})
	xorm.SetChecked(true)

	beegoOrm := widget.NewCheck("beegoOrm", func(bool) {})
	beegoOrm.SetChecked(true)

	jsonType := widget.NewSelect([]string{"仅生成", "生成并包含 omitempty"}, func(s string) {

	})
	jsonType.SetSelected("仅生成")

	json := widget.NewCheck("json", func(on bool) {
		if !on {
			jsonType.Hide()
		} else {
			jsonType.Show()
		}
	})
	json.SetChecked(true)

	excludeTables := widget.NewMultiLineEntry()
	excludeTables.SetPlaceHolder("多个数据表以回车换行")
	tryComplete := widget.NewRadio([]string{"是", "否"}, func(s string) {

	})
	tryComplete.Horizontal = true
	tryComplete.SetSelected("是")
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
