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
	"xorm.io/core"
)

func Screen(win fyne.Window, dbConf []interface{}) *fyne.Container {
	resultBox := widget.NewMultiLineEntry()
	resultBox.SetPlaceHolder(`准备就绪`)
	tables := widget.NewTabContainer()
	Sql2struct.Tables = []*core.Table{
		core.NewTable("article", nil),
		core.NewTable("category", nil),
	}
	for _, t := range Sql2struct.Tables {
		tableName := t.Name
		tableBox := widget.NewMultiLineEntry()
		tableBox.OnCursorChanged = func() {
			if rs, err := sql2struct(win, []string{tableName}, dbConf); err != nil {
				resultBox.SetText(err.Error())
			} else {
				resultBox.SetText(strings.ReplaceAll(string(rs), "\t", "    "))
				tableBox.SetText(tableName) //转换为表结构
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
