package spielbrett

import (
	"app/karten"
	"fmt"
	"math/rand"
	"os"
	"time"

	"fyne.io/fyne/v2"
)

type spielbrett struct {
	karten []karten.Karte
}

func New() *spielbrett {
	var sb *spielbrett
	sb = new(spielbrett)
	sb.karten = make([]karten.Karte, 12)

	inhaltDerKarten := beliebigerInhalt()

	for i := 0; i < 12; i++ {
		fmt.Println("", i)
		index := i
		k := karten.NewKarte(inhaltDerKarten[i],
			func() {
				sb.KarteAusgewaehlt(index)
			},
		)
		sb.karten[i] = k
	}

	return sb
}

// In der nachfolgenden Funktion wird ein Slice aus Strings mit 6
// Buchstaben/ Audiodateien gefüllt. Der Inhalt wird verdoppelt
// und im Anschluss zufällig angeordnet

func beliebigerInhalt() []string {

	// Namen der Autodateien werden ausgelesen und in files gespeichert.
	files, err := os.ReadDir("./AudiofilesMemoric")
	if err != nil {
		fmt.Println("Fataler Fehler!")
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}

	//	buchstaben := [6]string{"A", "B", "C", "D", "E", "F"}

	var verdoppelung []string

	for _, wert := range files {
		verdoppelung = append(verdoppelung, wert.Name(), wert.Name())
	}

	fmt.Println("Das ist das aktuelle Feld:", verdoppelung)

	// Initialisierung des Zufallswertes
	rand.Seed(time.Now().UnixNano())

	// Zufällige Anordnung der Werte im Feld.
	rand.Shuffle(len(verdoppelung), func(i, j int) {
		verdoppelung[i], verdoppelung[j] = verdoppelung[j], verdoppelung[i]
	})

	fmt.Println("Das ist das aktuelle Feld:", verdoppelung)

	return verdoppelung
}

func (sb *spielbrett) GetKartenFuerFyne() []fyne.CanvasObject {

	fyneKarten := make([]fyne.CanvasObject, len(sb.karten))
	for i, v := range sb.karten {
		fyneKarten[i] = fyne.CanvasObject(v)
	}
	return fyneKarten
}

func (sb *spielbrett) offeneKarten() int {
	anzahl := 0
	for _, v := range sb.karten {
		if v.IstOffen() {
			anzahl++
		}
	}
	return anzahl
}

func (sb *spielbrett) alleKartenSchliessen() {
	for _, v := range sb.karten {
		if v.IstOffen() {
			v.Schliessen()
		}
	}
}

func (sb *spielbrett) KarteAusgewaehlt(kartennr int) {
	fmt.Println("Karte wurde ausgewählt")
	fmt.Println("Kartennummer ist: ", kartennr)

	ausgewaehlteKarte := sb.karten[kartennr]

	//	ausgewaehlteKarte.Verschwinden()

	anzahlOffen := sb.offeneKarten()

	if anzahlOffen < 2 {
		ausgewaehlteKarte.Oeffnen()
	} else {
		if sb.offeneKartenGleich() {
			sb.offeneKartenVerschwinden()
		}
		sb.alleKartenSchliessen()
	}

}

// Funktion testet, ob die alle - hier zwei - geöffneten Karten gleichen Inhalt haben.
func (sb *spielbrett) offeneKartenGleich() bool {

	var ersterInhalt string

	for _, v := range sb.karten {
		if v.IstOffen() {
			if ersterInhalt == "" {
				ersterInhalt = v.Inhalt()
			} else if ersterInhalt != v.Inhalt() {
				return false
			}

		}
	}

	return true

}

// Die Funktion lässt alle geöffneten Karten verschwinden.
func (sb *spielbrett) offeneKartenVerschwinden() {

	for _, v := range sb.karten {
		if v.IstOffen() {
			v.Verschwinden()
		}
	}
}
