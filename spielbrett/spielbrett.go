package spielbrett

import (
	"fyne.io/fyne/v2"
)

type Spielbrett interface {
	GetBrett() fyne.CanvasObject
}
