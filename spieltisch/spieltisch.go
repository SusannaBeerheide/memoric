package spieltisch

import "fyne.io/fyne/v2"

// Vor.: keine
// Erg.: Eine neue Instanz des ADO Spieltisch ist geliefert.
// func New() *data // *data erfüllt das Interface spieltisch

type Spieltisch interface {
	// Vor.: keine
	// Erg.: Ein Canvasobjekt aus dem Fyne-Paket ist geliefert, welches den Spieltisch anzeigt.
	GetFyneTisch() fyne.CanvasObject
}
