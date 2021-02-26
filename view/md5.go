package view

import (
	"crypto/md5"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/atotto/clipboard"
	"log"
)

func NewMd5Container() *fyne.Container {
	content := container.NewVBox(
		widget.NewLabel("MD5 Convert"),
	)

	content.Add(widget.NewButton("Add more items", func() {
		content.Add(addMoreMd5ConvertItem())
	}))

	content.Add(addMoreMd5ConvertItem())

	return content
}

func addMoreMd5ConvertItem() *fyne.Container {
	input := widget.NewEntry()
	outputIn32L := widget.NewButton("NaN", nil)
	outputIn32H := widget.NewButton("NaN", nil)

	outputIn32L.OnTapped = func() {
		if err := clipboard.WriteAll(outputIn32L.Text); err != nil {
			log.Println(err)
		} else {
			notify()
		}
	}

	outputIn32H.OnTapped = func() {
		if err := clipboard.WriteAll(outputIn32H.Text); err != nil {
			log.Println(err)
		} else {
			notify()
		}
	}

	input.SetPlaceHolder("enter text")
	input.OnChanged = func(s string) {
		outputIn32L.SetText(fmt.Sprintf("%x", md5.Sum([]byte(s))))
		outputIn32H.SetText(fmt.Sprintf("%X", md5.Sum([]byte(s))))
	}

	return container.NewGridWithColumns(3,
		input,
		outputIn32L,
		outputIn32H,
	)
}
