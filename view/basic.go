package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strconv"
	"time"
)

const TIME_FORMAT = "2006-01-02 15:04:05"

var ct widget.Label
var cts widget.Label

func NewTimeContainer() *fyne.Container {
	content := container.NewVBox(
		widget.NewLabel("Current Time"),
		container.NewGridWithColumns(2, &ct, &cts),
		widget.NewLabel("Timestamp to Time"),
	)

	content.Add(widget.NewButton("Add more items", func() {
		content.Add(addMore())
	}))

	content.Add(addMore())

	return content
}

func RefreshTime() {
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ticker.C:
			setTime()
		}
	}
}

func addMore() *fyne.Container {
	output := widget.NewLabel("NaN")
	input := widget.NewEntry()

	input.SetPlaceHolder("enter timestamp")
	input.OnChanged = func(s string) {
		if ts, err := strconv.ParseInt(s, 10, 64); err != nil {
			output.SetText("NaN")
		} else {
			output.SetText(time.Unix(ts, 0).Format(TIME_FORMAT))
		}
	}

	return container.NewGridWithColumns(2,
		input,
		output,
	)
}

func setTime() {
	now := time.Now()
	ct.SetText(now.Format(TIME_FORMAT))
	cts.SetText(strconv.FormatInt(now.Unix(), 10))
}

func init() {
	setTime()
}
