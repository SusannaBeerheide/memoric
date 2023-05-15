package spielbrett

import (
	"app/karten"
	"fmt"

	"fyne.io/fyne/v2"
)

type spielbrett struct {
	karten []karten.Karte
}

func New() *spielbrett {
	var sb *spielbrett
	sb = new(spielbrett)
	sb.karten = make([]karten.Karte, 12)

	for i := 0; i < 12; i++ {
		fmt.Println("", i)
		index := i
		k := karten.NewKarte("memoric",
			func() {
				sb.KarteAusgewaehlt(index)
			},
		)
		sb.karten[i] = k
	}

	return sb
}

func (sb *spielbrett) GetKartenFuerFyne() []fyne.CanvasObject {

	fyneKarten := make([]fyne.CanvasObject, len(sb.karten))
	for i, v := range sb.karten {
		fyneKarten[i] = fyne.CanvasObject(v)
	}
	return fyneKarten
}

func (sb *spielbrett) KarteAusgewaehlt(kartennr int) {
	fmt.Println("Karte wurde ausgewählt")
	fmt.Println("Kartennummer ist: ", kartennr)
	//	var kartenNummerSpeicher []int = make([]int, 0)

	ausgewaehlteKarte := sb.karten[kartennr]
	//fmt.Println("Karte ist offen: ", ausgewaehlteKarte.IstOffen())

	if len(sb.kartenNummerSpeichern(kartennr)) == 0 {
		ausgewaehlteKarte.Oeffnen()
	} else if len(sb.kartenNummerSpeichern(kartennr)) == 1 {
		for _, w := range kartenNummerSpeicher {
			if w != kartennr {
				ausgewaehlteKarte.Oeffnen()
			}
		}

		// testen, ob bereits geöffnete Karten == angeklickte Karten sind
		// Falls Karte bereits geöffent wurde passiert nichts
		// Falls Karte noch nicht geöffnet wurde, wird die angeklickte Karte geöffnet
		// für später: Testen, ob Paar vorliegt nach öffnen der 2. Karte.
	} else if len(sb.kartenNummerSpeichern(kartennr)) == 2 {
		for _, w := range kartenNummerSpeicher { // Alle bereits geöffneten Karten werden geschlossen.
			zuSchließendeKarte := sb.karten[w]
			fmt.Println("Kartennummer der zu schließenden Karte lautet", w)
			zuSchließendeKarte.Schliessen()
		}
		ausgewaehlteKarte.Oeffnen() // Angeklickte Karte wird geöffnet.
	}

}

var kartenNummerSpeicher []int = make([]int, 0)

func (sb *spielbrett) kartenNummerSpeichern(kartennr int) []int {
	if len(kartenNummerSpeicher) < 2 {
		kartenNummerSpeicher = append(kartenNummerSpeicher, kartennr)
		fmt.Println("Inhalt vom karteNummerSpeicher lautet", kartenNummerSpeicher)
	} else {
		/*	for _, w := range kartenNummerSpeicher {
				zuSchließendeKarte := sb.karten[w]
				fmt.Println("Kartennummer der zu schließenden Karte lautet", w)
				zuSchließendeKarte.Schliessen()
			}
		*/
		kartenNummerSpeicher = nil
		kartenNummerSpeicher = append(kartenNummerSpeicher, kartennr)
		fmt.Println("Inhalt vom karteNummerSpeicher lautet", kartenNummerSpeicher, "die Länge ist", len(kartenNummerSpeicher))

	}

	return kartenNummerSpeicher

}
