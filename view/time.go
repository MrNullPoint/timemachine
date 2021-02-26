package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/atotto/clipboard"
	"log"
	"strconv"
	"time"
)

const TIME_FORMAT = "2006-01-02 15:04:05"

var ct *widget.Button
var cts *widget.Button

func NewTimeContainer() *fyne.Container {
	content := container.NewVBox(
		widget.NewLabel("Current Time"),
		container.NewGridWithColumns(2, ct, cts),
		widget.NewLabel("Timestamp to Time"),
	)

	content.Add(widget.NewButton("Add more items", func() {
		content.Add(addMoreTimeConvertItem())
	}))

	content.Add(addMoreTimeConvertItem())

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

func addMoreTimeConvertItem() *fyne.Container {
	input := widget.NewEntry()
	output := widget.NewButton("NaN", nil)

	output.OnTapped = func() {
		if err := clipboard.WriteAll(ct.Text); err != nil {
			log.Println(err)
		} else {
			notify()
		}
	}

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
	now := time.Now()

	ct = widget.NewButton(now.Format(TIME_FORMAT), nil)
	cts = widget.NewButton(strconv.FormatInt(now.Unix(), 10), nil)

	ct.OnTapped = func() {
		if err := clipboard.WriteAll(ct.Text); err != nil {
			log.Println(err)
		} else {
			notify()
		}
	}

	cts.OnTapped = func() {
		if err := clipboard.WriteAll(cts.Text); err != nil {
			log.Println(err)
		} else {
			notify()
		}
	}
}
