package sql2struct

import (
	"errors"
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"gormat/controllers/Sql2struct"
	"strings"
	"time"
)

func Screen(win fyne.Window) (c *fyne.Container) {
	result := widget.NewMultiLineEntry()
	c = &fyne.Container{}
	if err := Sql2struct.InitDb(); err != nil {
		return
	}
	ts, err := Sql2struct.DBMetas(nil, Sql2struct.Configs().ExcludeTables, Sql2struct.Configs().TryComplete)
	if err != nil {
		return
	}
	var tables []fyne.CanvasObject
	for _, v := range ts {
		tName := v.Name
		tables = append(tables, widget.NewCheck(tName, func(checked bool) {
			if checked {
				if rs, err := sql2struct(win, []string{tName}); err != nil {
					result.SetText(err.Error())
				} else {
					result.SetText(strings.ReplaceAll(string(rs), "\t", "    "))
				}
			}
		}))
	}
	c = fyne.NewContainerWithLayout(
		layout.NewGridLayoutWithRows(1),
		widget.NewGroupWithScroller("选择表",
			tables...,
		),
		widget.NewScrollContainer(result),
	)
	return
}

func sql2struct(win fyne.Window, ts []string) (result []byte, err error) {
	status := dialog.NewProgress("转换中...", strings.Join(ts, "\n"), win)
	go func() {
		num := 0.0
		for num < 1.0 {
			time.Sleep(50 * time.Millisecond)
			status.SetValue(num)
			num += 0.01
		}
		status.SetValue(1)
	}()
	status.Show()
	if result, err = Sql2struct.NewGenTool().Gen(ts); err != nil {
		dialog.ShowError(errors.New("转换失败"), win)
		return
	}
	status.Hide()
	return
}
