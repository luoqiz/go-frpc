package utils

import "fyne.io/fyne"

func SendNotifiction(content string) {
	fyne.CurrentApp().SendNotification(&fyne.Notification{
		Title:   "go-frpc 客户端",
		Content: content,
	})
}
