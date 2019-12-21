package _app

import (
	"errors"
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"gormat/common"
	"strings"
	"time"
)

func Sql2struct(win fyne.Window, ts []string) (str string, err error) {
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
	if str, err = common.NewGenTool().Gen(ts); err != nil {
		dialog.ShowError(errors.New("转换失败"), win)
		return
	}
	status.Hide()
	return
}
