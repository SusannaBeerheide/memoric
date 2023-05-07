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

	var karten []spielkarte.KarteImpl

	for i := 0; i < 12; i++ {
		karten = append(karten, spielkarte.NewKarte())
	}

	karte := spielkarte.NewKarte()

	w.SetContent(karte.GetQuadrat())

	//w.SetContent(widget.NewLabel("Hello World!"))
	w.ShowAndRun()
}
