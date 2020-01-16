package sql2struct

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/liudanking/gorm2sql/program"
	"go/ast"
	"go/parser"
	"go/token"
	"gormat/app/config"
	"gormat/internal/Sql2struct"
	"gormat/internal/Sql2struct/quickly"
	"gormat/internal/Sql2struct/sqlorm"
	"path/filepath"
	"strings"
	"time"
	"xorm.io/core"
)

func Screen(win fyne.Window, dbConf []interface{}) *fyne.Container {
	if err := Sql2struct.InitDb(dbConf); err != nil {
		return fyne.NewContainerWithLayout(
			layout.NewGridLayout(1),
			widget.NewLabel(err.Error()),
		)
	}
	resultBox := widget.NewMultiLineEntry()
	resultBox.SetPlaceHolder(`准备就绪`)
	var tables = widget.NewTabContainer()
	var currentTable = make(chan *widget.TabItem)
	if tbs, err := Sql2struct.DBMetas(nil, Sql2struct.Configs().ExcludeTables, Sql2struct.Configs().TryComplete); err == nil {
		for _, t := range tbs {
			tables.Items = append(tables.Items, widget.NewTabItemWithIcon(t.Name, config.Table, widget.NewMultiLineEntry()))
		}
		go func(dbConf []interface{}) {
			for {
				time.Sleep(time.Microsecond * 50)
				if <-currentTable != tables.CurrentTab() {
					var currentT *core.Table
					for _, t := range tbs {
						if t.Name == tables.CurrentTab().Text {
							currentT = t
						}
					}
					if currentT == nil {
						return
					}
					_ = Sql2struct.InitDb(dbConf)
					tableBox := widget.NewMultiLineEntry()
					if result, err := Sql2struct.NewGenTool().Gen([]string{currentT.Name}, dbConf); err != nil {
						resultBox.SetText(``)
						tableBox.SetText(err.Error())
					} else {
						resultBox.SetText(strings.ReplaceAll(string(result), "\t", "    "))
						url := strings.Split(currentT.Name, "_")
						for k, v := range url {
							url[k] = strings.Title(v)
						}
						if file, err := parser.ParseFile(
							token.NewFileSet(),
							fmt.Sprintf(currentT.Name+"_%d", time.Now().Nanosecond()),
							resultBox.Text, parser.ParseComments); err != nil {
							tableBox.SetText(fmt.Sprintf("generate sql failed:%v", err))
						} else {
							var sqlStr []string
							types := program.FindMatchStruct([]*ast.File{file}, func(structName string) bool {
								match, _ := filepath.Match(strings.Join(url, ""), structName)
								return match
							})
							for _, typ := range types {
								ms, err := sqlorm.NewSqlGenerator(typ)
								if err != nil {
									tableBox.SetText(fmt.Sprintf("create model struct failed:%v", err))
									continue
								} else {
									if sql, err := ms.GetCreateTableSql(currentT); err != nil {
										tableBox.SetText(fmt.Sprintf("generate sql failed:%v", err))
									} else {
										sqlStr = append(sqlStr, sql)
									}
								}
							}
							tableBox.SetText(strings.Join(sqlStr, "\n\n"))
						}
					}
					tables.CurrentTab().Content = tableBox
					tables.Refresh()
				}
			}
		}(dbConf)
	} else {
		return fyne.NewContainerWithLayout(
			layout.NewGridLayout(1),
			widget.NewLabel(err.Error()),
		)
	}
	tableBox := widget.NewGroupWithScroller("表结构")
	if len(tables.Items) > 0 {
		go func() {
			currentTable <- &widget.TabItem{}
			for {
				currentTable <- tables.CurrentTab()
			}
		}()
		tables.SetTabLocation(widget.TabLocationLeading)
		tableBox.Append(tables)
	}
	return fyne.NewContainerWithLayout(
		layout.NewGridLayout(1),
		tableBox,
		widget.NewScrollContainer(resultBox),
	)
}

func QuickScreen() *fyne.Container {
	result := widget.NewMultiLineEntry()
	data := widget.NewMultiLineEntry()
	data.OnChanged = func(s string) {
		if s == "" {
			result.SetText(s)
			return
		}
		s = strings.ReplaceAll(s, "`", "")
		blocks, err := quickly.MatchStmt(strings.NewReader(s))
		if err != nil {
			result.SetText(err.Error())
			return
		}
		r := ""
		for _, v := range blocks {
			ss, err := quickly.HandleStmtBlock(v).GenType()
			if err != nil {
				r += err.Error()
				continue
			}
			r += string(ss)
		}
		result.SetText(strings.ReplaceAll(r, "\t", "    "))
	}
	data.PlaceHolder = `CREATE TABLE ` + "`" + `your_struct` + "`" + ` (
  ` + "`" + `id` + "`" + ` int(11) NOT NULL AUTO_INCREMENT,
);`
	result.SetPlaceHolder(`type YourStruct struct {
    ID uint ` + "`" + `gorm:"column:id;not null;AUTO_INCREMENT;type:int(11)" json:"id,omitempty"` + "`" + `
}`)
	return fyne.NewContainerWithLayout(
		layout.NewGridLayoutWithColumns(1),
		widget.NewScrollContainer(data),
		widget.NewScrollContainer(result),
	)
}
