/*
@Time : 2019/12/20 16:04
@Software: GoLand
@File : welcome
@Author : Bingo <airplayx@gmail.com>
*/
package _app

import (
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"net/url"
)

func WelcomeScreen() fyne.CanvasObject {
	logo := canvas.NewImageFromResource(fyne.NewStaticResource("logo", Logo))
	logo.SetMinSize(fyne.NewSize(275, 470))

	airplayx, _ := url.Parse("http://github.com/airplayx")
	fyneIo, _ := url.Parse("http://fyne.io/fyne")

	return widget.NewVBox(
		widget.NewLabelWithStyle("", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		layout.NewSpacer(),
		widget.NewHBox(layout.NewSpacer(), logo, layout.NewSpacer()),
		widget.NewHyperlinkWithStyle("github.com/airplayx", airplayx, fyne.TextAlignCenter, fyne.TextStyle{}),
		widget.NewHyperlinkWithStyle("fyne.io/fyne", fyneIo, fyne.TextAlignCenter, fyne.TextStyle{}),
		layout.NewSpacer(),
	)
}
