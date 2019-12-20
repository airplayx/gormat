/*
@Time : 2019/12/20 16:36
@Software: GoLand
@File : special_tab
@Author : Bingo <airplayx@gmail.com>
*/
package tab

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

func Special() fyne.Widget {
	return &widget.Form{
		OnCancel: func() {

		},
		OnSubmit: func() {

		},
		Items: []*widget.FormItem{},
	}
}
