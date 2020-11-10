package service

import (
	"fyne.io/fyne"
	"fyne.io/fyne/container"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
	"go-frpc/frp"
	"go-frpc/utils"
)

func InfoScreen(w fyne.Window) fyne.CanvasObject {

	proxyLable := widget.NewLabel("请输入http代理")
	proxyEntry := widget.NewEntry()
	proxyBox := container.NewHBox(proxyLable, proxyEntry)
	// 文件下载按钮
	downloadButton := widget.NewButton("download",
		func() {
			if proxyEntry.Text != "" {
				utils.AppConfig.ProxyAddr = proxyEntry.Text
			}
			prog := dialog.NewProgress("文件下载", "下载进度", w)
			prog.Show()
			frp.Download(func(length, downLen int64) {
				proc := float64(downLen)/float64(length) + 0.005
				println(downLen, length, proc)
				if proc < 1.0 {
					prog.SetValue(proc)
				} else {
					prog.SetValue(1)
					prog.Hide()
				}
			})

		},
	)
	return container.NewVBox(
		proxyBox,
		downloadButton)
}
