package main

import (
	"app/spieltisch"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	//	"fyne.io/fyne/v2/widget"
)

func main() {
	// Anlegen eines neuen Fyne-Fensters
	a := app.New()
	os.Setenv("FYNE_SCALE", "1.3")
	w := a.NewWindow("memoric")
	// Festlegen der Fensterausgangsgröße
	w.Resize(fyne.NewSize(800, 600))

	// Anlegen der Variable für den Spieltisch und füllen mit dem neu erschaffenen ADO Spieltisch
	spieltisch := spieltisch.New()

	// Der Spieltich wird als Fensterinhalt gesetzt.
	w.SetContent(spieltisch.GetFyneTisch())

	//w.SetContent(widget.NewLabel("Hello World!"))

	// Das Fenster wird angezeigt und dessen Inhalte laufen.
	w.ShowAndRun()
}
