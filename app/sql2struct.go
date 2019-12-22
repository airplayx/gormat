package _app

import (
	"errors"
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"gormat/common"
	"strings"
	"time"
)

func Sql2structScreen(win fyne.Window, result *widget.Entry) *fyne.Container {
	if err := common.LoadConfig(); err != nil {
		dialog.ShowError(errors.New("加载配置文件错误"), win)
		return nil
	}
	if err := common.InitDb(); err != nil {
		dialog.ShowError(errors.New("数据库连接失败"), win)
		return nil
	}
	ts, err := common.DBMetas(nil, common.Configs().ExcludeTables, common.Configs().TryComplete)
	if err != nil {
		dialog.ShowError(errors.New("数据库信息获取失败"), win)
		return nil
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
	return fyne.NewContainerWithLayout(
		layout.NewGridLayoutWithRows(1),
		//左侧表数据
		widget.NewGroupWithScroller("选择表",
			tables...,
		),
		//右侧结构体
		widget.NewScrollContainer(result),
	)
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
	if err = common.LoadConfig(); err != nil {
		dialog.ShowError(errors.New("加载配置文件错误"), win)
		return
	}
	if result, err = common.NewGenTool().Gen(ts); err != nil {
		dialog.ShowError(errors.New("转换失败"), win)
		return
	}
	status.Hide()
	return
}
