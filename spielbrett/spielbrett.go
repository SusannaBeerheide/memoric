package spielbrett

import (
	"fyne.io/fyne/v2"
)

// Vor.: Es existeren die Klassen "karten" und "score".
// Erg.: Eine neue Instanz des ADO Spielbrett ist geliefert.
// func New(score score.Score) *spielbrett // *data erfüllt das Interface spielbrett

type Spielbrett interface {
	// Vor.: keine
	// Erg.: Ein Canvasobjekt aus dem Fyne-Paket ist geliefert, welches das Spielbrett anzeigt.
	GetBrett() fyne.CanvasObject
}
