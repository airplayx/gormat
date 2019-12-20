/*
@Time : 2019/12/20 16:40
@Software: GoLand
@File : option
@Author : Bingo <airplayx@gmail.com>
*/
package tab

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

func Option() fyne.Widget {
	targetDir := widget.NewEntry()
	targetDir.SetPlaceHolder("./models")
	gorm := widget.NewCheck("gorm", func(bool) {})
	gorm.SetChecked(true)
	gorm.Hide()

	json := widget.NewSelect([]string{"仅生成", "生成并包含 omitempty", "不生成"}, func(s string) {

	})
	json.SetSelected("仅生成")
	excludeTables := widget.NewMultiLineEntry()
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
			{Text: "json标签", Widget: json},
			{Text: "排除表", Widget: excludeTables},
			{Text: "始终执行", Widget: tryComplete},
		},
	}
}
