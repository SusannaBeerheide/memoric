package score

import "fyne.io/fyne/v2/data/binding"

type Score interface {
	GetAktuellerSpieler() bool
	Spieler1() binding.String
	Spieler2() binding.String
	PaarGefunden()
	SpielerWechseln()
}
