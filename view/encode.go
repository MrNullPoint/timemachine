package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"net/url"
)

func NewEncodeContainer() *fyne.Container {
	content := container.NewVBox()

	content.Add(widget.NewLabel("URL Encode"))

	encodeOutput := widget.NewEntry()
	encodeOutput.MultiLine = true

	encodeInput := widget.NewEntry()
	encodeInput.MultiLine = true
	encodeInput.SetPlaceHolder("Enter text for url encode...")

	encodeInput.OnChanged = func(s string) {
		encodeOutput.SetText(url.QueryEscape(s))
	}

	content.Add(container.NewGridWithColumns(2,
		encodeInput,
		encodeOutput,
	))

	content.Add(widget.NewLabel("URL Decode"))

	decodeOutput := widget.NewEntry()
	decodeOutput.MultiLine = true

	decodeInput := widget.NewEntry()
	decodeInput.MultiLine = true
	decodeInput.SetPlaceHolder("Enter text for url decode...")

	decodeInput.OnChanged = func(s string) {
		decodeStr, _ := url.QueryUnescape(s)
		decodeOutput.SetText(decodeStr)
	}

	content.Add(container.NewGridWithColumns(2,
		decodeInput,
		decodeOutput,
	))

	return content
}
