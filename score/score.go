package score

import "fyne.io/fyne/v2/data/binding"

// Vor.: keine
// Erg.: Eine neue Instanz des ADT Score ist geliefert.
// func New() *data // // *data erfüllt das Interface score

type Score interface {
	// Vor.: Es gibt einen aktuellen Spieler.
	// Erg.: Der aktuelle Spieler ist geliefert, wobei false = Spieler 1 ist dran, true = Spieler 2 ist dran, bedeutet.
	GetAktuellerSpieler() bool

	// Vor.: Der Spieler 1 besitzt einen Punktestand.
	// Erg.: Der Spielstand von Spieler 1 als String im Fyne-Paket ist geliefert.
	Spieler1() binding.String

	// Vor.: Der Spieler 1 besitzt einen Punktestand.
	// Erg.: Der Spielstand von Spieler 2 als String im Fyne-Paket ist geliefert.
	Spieler2() binding.String

	// Vor.: Es gibt einen aktuellen Spieler.
	// Erg.: Der aktuelle Spieler, d.h. der Spieler der gerade zieht, ist als String im Fyne-Paket geliefert.
	WerDranIst() binding.String

	// Vor.: Beide Spieler besitzen einen Punktestand.
	// Eff.: Der Spielstand des aufrufenden Spielers ist um 1 erhöht und die Spielstände beider Spieler sind gesetzt.
	PaarGefunden()

	// Vor.: Es gibt einen aktuellen Spieler.
	// Eff.: Der aktuelle Spieler ist der zuvor nicht aktuelle Spieler.
	SpielerWechseln()
}
