package sql2struct

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	_app "gormat/app"
	"gormat/controllers/Sql2struct"
	"strings"
)

func Screen(win fyne.Window, dbConf []interface{}) *fyne.Container {
	if err := Sql2struct.InitDb(dbConf); err != nil {
		return &fyne.Container{}
	}
	resultBox := widget.NewMultiLineEntry()
	resultBox.SetPlaceHolder(`准备就绪`)
	tables := widget.NewTabContainer()
	if tbs, err := Sql2struct.DBMetas(
		nil, Sql2struct.Configs().ExcludeTables, Sql2struct.Configs().TryComplete); err == nil {
		for _, t := range tbs {
			tableName := t.Name
			tableItem := widget.NewMultiLineEntry()
			tableItem.OnCursorChanged = func() {
				if result, err := Sql2struct.NewGenTool().Gen([]string{tableName}, dbConf); err != nil {
					resultBox.SetText(err.Error())
				} else {
					resultBox.SetText(strings.ReplaceAll(string(result), "\t", "    "))
					tableItem.SetText(tableName) //转换为表结构
				}
				resultBox.Refresh()
			}
			tables.Items = append(tables.Items, widget.NewTabItemWithIcon(t.Name, _app.Table, tableItem))
		}
	} else {
		return &fyne.Container{}
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
