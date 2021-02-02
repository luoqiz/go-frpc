package service

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"go-frpc/src/frp"
)

func StatusScreen(_ fyne.Window) fyne.CanvasObject {

	pid := frp.CheckStatus()
	showText := ""
	if pid != 0 {
		showText = fmt.Sprintln("running ... pid:", pid)
	} else {
		showText = fmt.Sprintln("frp not runing...")
	}
	statusShow := widget.NewLabelWithStyle(showText, fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	// 启动客户按钮
	startButton := widget.NewButton("start", func() {
		pid := frp.Start()
		if pid != 0 {
			statusShow.SetText(fmt.Sprintln("running pid: ", pid))
		} else {
			statusShow.SetText(fmt.Sprintln("please download frp... ", pid))
		}
	})

	// 关闭客户端按钮
	stopButton := widget.NewButton("stop", func() {
		if frp.CheckStatus() != 0 {
			frp.Stop()
			statusShow.SetText("frp success closed...")
		} else {
			statusShow.SetText("frp not running...")
		}

	})

	return container.NewVBox(
		container.New(layout.NewGridLayout(2),
			startButton,
			stopButton,
		),
		statusShow,
	)
}
