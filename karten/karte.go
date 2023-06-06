package karten

import "fyne.io/fyne/v2"

// Vor.:
// Erg.: Eine neue Instanz des ADT Karte ist geliefert.
// func NewKarte(text string, brettFunc func()) *data // // *data erfüllt das Interface karte

type Karte interface {
	fyne.CanvasObject

	//Vor.: Eine Karte wurde  gerade angeklickt.
	//Eff.: Die mittels onTapped übergebene Funktion wird ausgeführt.
	Tapped(_ *fyne.PointEvent)

	//Vor.: Keine.
	//Eff.: Die interne Funktion newKarteRenderer wird aufgerufen (nötig für Fyne Funktionalitäten).
	//Erg.: Die grafische Darstellung der Karte mittels interner Fyne Funktionalität wird realisiert.
	CreateRenderer() fyne.WidgetRenderer

	//Vor.: Keine.
	//Erg.: true ist geliefert, wenn die Karte geöffnet ist, anderernfalls ist false geliefert.
	IstOffen() bool

	//Vor.: Die Karte wurde angeklickt.
	//Eff.: Die hinterlegte Musik wird abgespielt und die Refresh-Funktion wird ausgeführt (deren Eff. hier: Kartenfarbe blau).
	//Erg.: w.offen ist true.
	Oeffnen()

	//Vor.: Eine geöffnete Karte lag vor.
	//Eff.: Das Abspielen der Musik wird gestoppt.
	MusikStoppen()

	//Vor.: w.offen war true.
	//Eff.: Die Musik wird gestoppt und die Refresh-Funktion wird ausgeführt (deren Eff. hier: Kartenfarbe wieder rot).
	//Erg.: w.offen ist false.
	Schliessen()

	//Vor.: w.offen war true.
	//Eff.: Die Refresh-Funktion wird ausgeführt (deren Eff. hier: Kartenfarbe transparent).
	//Erg.: w.weg ist true und w.offen ist false.
	Verschwinden()

	//Vor.: Keine.
	//Erg.: Gibt den Text der Spielkarte zurück, z.B. "memoric".
	Inhalt() string
}
