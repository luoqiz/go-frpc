package service

import (
	"fyne.io/fyne"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
	"go-frpc/frp"
)

func AdvancedScreen(_ fyne.Window) fyne.CanvasObject {
	// 文本框
	content := widget.NewMultiLineEntry()
	// 保存按钮
	saveButton := widget.NewButton("save", func() {
		frp.SetContent(content.Text)
	})
	// 重新加载按钮
	reloadButton := widget.NewButton("reload", func() {
		fc, err := frp.FullContent()
		if err != nil {
			content.SetText("not fount the file : frpc.ini")
		} else {
			content.SetText(fc)
		}

	})
	buttons := container.NewGridWithColumns(2, saveButton, reloadButton)
	scroll := container.NewScroll(content)
	scroll.SetMinSize(fyne.NewSize(400, 500))
	return container.NewMax(container.NewVBox(buttons, container.NewAdaptiveGrid(1, scroll)))
}
