package sql2struct

import (
	"errors"
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	_app "gormat/app"
	"gormat/controllers/Sql2struct"
	"strings"
)

func Screen(win fyne.Window, dbConf []interface{}) *fyne.Container {
	resultBox := widget.NewMultiLineEntry()
	tables := widget.NewTabContainer()
	for _, t := range Sql2struct.Tables {
		tableName := t.Name
		tableBox := widget.NewMultiLineEntry()
		tableBox.SetText(tableName) //读取表结构
		tableBox.OnCursorChanged = func() {
			if rs, err := sql2struct(win, []string{tableName}, dbConf); err != nil {
				resultBox.SetText(err.Error())
			} else {
				resultBox.SetText(strings.ReplaceAll(string(rs), "\t", "    "))
			}
			resultBox.Refresh()
		}
		tables.Items = append(tables.Items, widget.NewTabItemWithIcon(t.Name, _app.Table, tableBox))
	}
	tableBox := widget.NewGroupWithScroller("表结构")
	if len(tables.Items) > 0 {
		tables.SetTabLocation(widget.TabLocationLeading)
		tableBox.Append(tables)
	}
	return fyne.NewContainerWithLayout(
		layout.NewGridLayout(1),
		tableBox,
		widget.NewScrollContainer(resultBox),
	)
}

func sql2struct(win fyne.Window, ts []string, dbConf []interface{}) (result []byte, err error) {
	if result, err = Sql2struct.NewGenTool().Gen(ts, dbConf); err != nil {
		dialog.ShowError(errors.New(err.Error()), win)
		return
	}
	return
}
