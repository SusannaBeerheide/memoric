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
	w := a.NewWindow("Rectangle")
	w.Resize(fyne.NewSize(800, 800))

	memorybrett := spielbrett.New()

	grid := container.New(layout.NewGridLayout(4), memorybrett.GetFyneKarten()...)

	w.SetContent(grid)

	w.ShowAndRun()
}
