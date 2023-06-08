package karten

import "fyne.io/fyne/v2"

// Vor.: Es existiert ein Score,welcher als Text
// Erg.: Eine neue Instanz des ADT Karte ist geliefert.
// func NewKarte(text string, brettFunc func()) *data // // *data erfüllt das Interface karte

type Karte interface {
	//Vor.: Eine Karte wurde  gerade angeklickt.
	//Eff.: Die mittels onTapped übergebene Funktion wird ausgeführt.
	Tapped(_ *fyne.PointEvent)

	//Vor.: Keine.
	//Erg.: true ist geliefert, wenn die Karte geöffnet ist, anderernfalls ist false geliefert.
	IstOffen() bool

	//Vor.: Die Karte wurde angeklickt.
	//Eff.: w.offen ist true. Die hinterlegte Musik wird abgespielt und die Refresh-Funktion wird ausgeführt (deren Eff. hier: Kartenfarbe blau).
	Oeffnen()

	//Vor.: w.offen war true.
	//Eff.: w.offen ist false. Die Musik wird gestoppt und die Refresh-Funktion wird ausgeführt (deren Eff. hier: Kartenfarbe wieder rot).
	Schliessen()

	//Vor.: Eine geöffnete Karte lag vor.
	//Eff.: Das Abspielen der Musik wird gestoppt.
	MusikStoppen()

	//Vor.: w.offen war true.
	//Eff.: w.weg ist true und w.offen ist false. Die Refresh-Funktion wird ausgeführt (deren Eff. hier: Kartenfarbe transparent).
	Verschwinden()

	//Vor.: Keine.
	//Erg.: Der Text der Spielkarte - z.B. "memoric" - ist geliefert.
	Inhalt() string

	// siehe Dokumentation im Fyne-Paket
	fyne.CanvasObject

	//Vor.: Keine.
	//Eff.: Die interne Funktion newKarteRenderer wird aufgerufen (nötig für Fyne Funktionalitäten).
	//Erg.: Die grafische Darstellung der Karte mittels interner Fyne Funktionalität wird realisiert.
	CreateRenderer() fyne.WidgetRenderer
}
