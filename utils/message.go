package utils

import "fyne.io/fyne"

func SendNotifiction(content string) {
	fyne.CurrentApp().SendNotification(&fyne.Notification{
		Title:   "frpc",
		Content: content,
	})
}
