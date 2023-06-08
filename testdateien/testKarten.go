package main

import (
	"app/karten"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	// Anlegen eines neuen Fyne-Fensters
	a := app.New()
	w := a.NewWindow("Kartentest")
	// Festlegen der Fensterausgangsgröße
	w.Resize(fyne.NewSize(200, 200))

	// Anlegen der Variable für den Spieltisch und füllen mit dem neu erschaffenen ADO Spieltisch
	karte := karten.NewKarte("kein Pfad",
		func() {
			fmt.Println("Yo")
		},
	)

	// Der Spieltich wird als Fensterinhalt gesetzt.
	w.SetContent(karte)

	// Das Fenster wird angezeigt und dessen Inhalte laufen.
	w.ShowAndRun()
}
