package spielbrett

import "app/karten"

type Spielbrett interface {
	GetKarten() []*karten.Karte
}
