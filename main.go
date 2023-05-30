package main

import (
	"app/spielbrett"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	//	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("memoric")
	w.Resize(fyne.NewSize(800, 800))

	memorybrett := spielbrett.New()

	grid := container.New(layout.NewGridLayout(4),
		memorybrett.GetKartenFuerFyne()...,
	// karten.NewKarte("Testing again", func() {}),
	)

	w.SetContent(grid)

	//w.SetContent(widget.NewLabel("Hello World!"))
	w.ShowAndRun()
}
