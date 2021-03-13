package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func NewMainView() fyne.CanvasObject {
	timeItem := container.NewTabItem("time", NewTimeContainer())
	md5Item := container.NewTabItem("md5", NewMd5Container())
	encodeItem := container.NewTabItem("encode", NewEncodeContainer())
	content := container.NewAppTabs(timeItem, md5Item, encodeItem)
	return content
}

func notify() {
	fyne.CurrentApp().SendNotification(&fyne.Notification{
		Title:   "Time Machine",
		Content: "Data has copied...",
	})
}
