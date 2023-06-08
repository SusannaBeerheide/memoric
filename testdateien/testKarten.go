package main

import (

	hier müsste noch das Paket Karten import werden

	"fyne.io/fyne"
	"fyne.io/fyne/app"
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
			???
		}
	)

	// Der Spieltich wird als Fensterinhalt gesetzt.
	w.SetContent(karten. ??? )


	// Das Fenster wird angezeigt und dessen Inhalte laufen.
	w.ShowAndRun()
}
