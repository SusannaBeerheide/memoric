package main

import (
	"app/spielkarte"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	//	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Rectangle")
	w.Resize(fyne.NewSize(800, 800))

	karte := new(spielkarte.KarteImpl)
	w.SetContent(karte.quadrat)

	//w.SetContent(widget.NewLabel("Hello World!"))
	w.ShowAndRun()
}
