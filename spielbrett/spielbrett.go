package spielbrett

import (
	"app/spielkarten"

	"fyne.io/fyne/v2"
)

type Spielbrett interface {
	GetKarten() []spielkarten.Karte
	GetFyneKarten() []fyne.CanvasObject
}
