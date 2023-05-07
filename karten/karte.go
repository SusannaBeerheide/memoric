package karten

import "fyne.io/fyne/v2"

type Karte interface {
	fyne.CanvasObject

	CreateRenderer() fyne.WidgetRenderer

	IstOffen() bool

	Oeffnen()

	Schliessen()
}
