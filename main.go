package main

import (
	"app/spieltisch"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	//	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("memoric")
	w.Resize(fyne.NewSize(800, 800))

	spieltisch := spieltisch.New()

	w.SetContent(spieltisch.GetFyneTisch())

	//w.SetContent(widget.NewLabel("Hello World!"))
	w.ShowAndRun()
}
