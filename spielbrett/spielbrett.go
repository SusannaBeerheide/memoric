package spielbrett

import (
	"fyne.io/fyne/v2"
)

type Spielbrett interface {
	GetKarten() []fyne.CanvasObject
}
