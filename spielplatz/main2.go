/*package main

import (
	"image/color"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Grid Layout")

	text1 := canvas.NewText("1", color.Black)
	text2 := canvas.NewText("2", color.Black)
	text3 := canvas.NewText("3", color.Black)
	text4 := canvas.NewText("4", color.Black)
	text5 := canvas.NewText("5", color.Black)
	text6 := canvas.NewText("6", color.Black)
	grid := container.New(layout.NewGridLayout(2), text1, text2, text3, text4, text5, text6)
	myWindow.SetContent(grid)
	myWindow.ShowAndRun()
}
*/
package main2

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Form Layout")

	label1 := widget.NewLabel("Label 1")
	value1 := widget.NewLabel("Value das ist jetzt ganz lang und noch länger")
	label2 := widget.NewLabel("Label 2")
	value2 := widget.NewLabel("Something")
	grid := container.New(layout.NewGridLayout(3), label1, value1, label2, value2)
	myWindow.SetContent(grid)
	myWindow.ShowAndRun()
}
