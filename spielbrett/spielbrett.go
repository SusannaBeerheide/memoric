package spielbrett

import "app/karten"

type Spielbrett interface {
	GetKartenFuerFyne() []*karten.Karte
}
