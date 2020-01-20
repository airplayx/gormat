/*
@Time : 2020/1/16 12:08
@Software: GoLand
@File : NewTabContainer_test
@Author : Bingo <airplayx@gmail.com>
*/
package test

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"log"
	"testing"
	"time"
)

func Test_tabContainer(t *testing.T) {
	w := app.NewWithID("Gopher").NewWindow("test")
	setting := widget.NewTabContainer(
		widget.NewTabItem("aaaaaa", fyne.NewContainer()),
		widget.NewTabItem("bbbbbb", fyne.NewContainer()),
		widget.NewTabItem("cccccc", fyne.NewContainer()),
	)
	current := make(chan string)
	go func() {
		for true {
			current <- setting.CurrentTab().Text
		}
	}()
	go func() {
		for true {
			time.Sleep(time.Microsecond * 200)
			if <-current != setting.CurrentTab().Text {
				fmt.Println(<-current)
			}
		}
	}()
	w.SetContent(setting)
	w.Resize(fyne.Size{Width: 650, Height: 300})
	w.CenterOnScreen()
	w.ShowAndRun()
}

func Test_AddContainer(t *testing.T) {
	w := app.NewWithID("Gopher").NewWindow("test")
	setting := widget.NewTabContainer(
		widget.NewTabItem("aaaaaa", fyne.NewContainer()),
		widget.NewTabItem("cccccc", fyne.NewContainer()),
	)
	setting.Append(widget.NewTabItem("ddd",
		fyne.NewContainer(widget.NewLabel("addddddd444444d"), widget.NewButton("adddddddd", func() {
			setting.Append(widget.NewTabItem("cccccc", fyne.NewContainer()))
			log.Println(`add success!`)
		}), widget.NewButton("adddddddd", func() {
			setting.Append(widget.NewTabItem("cccccc", fyne.NewContainer()))
			log.Println(`add success!`)
		}))))
	w.SetContent(setting)
	w.Resize(fyne.Size{Width: 650, Height: 300})
	w.CenterOnScreen()
	w.ShowAndRun()
}
