package service

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"go-frpc/src/frp"
	"go-frpc/src/utils"
)

func InfoScreen(w fyne.Window) fyne.CanvasObject {

	proxyEntry := widget.NewEntry()
	proxyBox := widget.NewForm(widget.NewFormItem("http代理", proxyEntry))
	// 文件下载按钮
	downloadButton := widget.NewButton("download && update",
		func() {

			// 若是填写了代理则将代理设置到软件全局
			if proxyEntry.Text != "" {
				utils.AppConfig.ProxyAddr = proxyEntry.Text
			}
			prog := dialog.NewProgress("文件下载", "下载进度", w)
			prog.Show()
			frp.Download(func(length, downLen int64) {
				proc := float64(downLen)/float64(length) + 0.005
				if proc < 1.0 {
					prog.SetValue(proc)
				} else {
					prog.SetValue(1)
					prog.Hide()
				}
			})

			// 检测是否已存在配置文件，不存在则创建
			if frp.GetIniFilePath() == "" {
				utils.CreateFile("./frpc.ini")
			}
		},
	)
	return container.NewVBox(
		proxyBox,
		downloadButton)
}
