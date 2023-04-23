package main

import (
	"image/color"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
//	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	a := app.New()
	w := a.NewWindow("Rectangle")
	w.Resize(fyne.NewSize(800, 800))

	

	rect := canvas.NewRectangle(color.Black)
	w.SetContent(rect)

	//w.SetContent(widget.NewLabel("Hello World!"))
	w.ShowAndRun()
}