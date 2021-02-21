package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/MrNullPoint/timemachine/view"
)

func main() {
	go view.RefreshTime()

	tApp := app.New()

	tWindow := tApp.NewWindow("Time Machine")

	tWindow.SetContent(view.NewTimeContainer())

	tWindow.ShowAndRun()
}
