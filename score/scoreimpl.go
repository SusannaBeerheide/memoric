package score

import (
	"strconv"

	"fyne.io/fyne/v2/data/binding"
)

type data struct { // Der Spielstand/ Score besteht aus:
	aktuellerSpieler bool           // der Angabe des aktuellen Spielers, wobei false = Spieler 1 ist dran, true = Spieler 2 ist dran bedeutet
	scoreSpieler1    int            // dem Spielstand von Spieler 1, welcher mit 0 initialisiert ist
	scoreSpieler2    int            // dem Spielstand von Spieler 2, welcher mit 0 initialisiert ist
	spieler1String   binding.String // der Spielstand des Spielers 1 als binding.String
	spieler2String   binding.String // der Spielstand des Spielers 2 als binding.String
	dranIst          binding.String // zur Anzeige, wer dran ist.
}

func New() *data {
	var sc *data
	sc = new(data)
	sc.spieler1String = binding.NewString() // Erzeugen einer neuen Instanz ADT binding.String mit dem Fyne-Paket
	sc.spieler2String = binding.NewString() // Erzeugen einer neuen Instanz ADT binding.String mit dem Fyne-Paket
	sc.setScoreStrings()                    // Setzt die aktuellen Spielstände der zwei Spieler Strings um.
	sc.dranIst = binding.NewString()        // Erzeugen einer neuen Instanz ADT binding.String mit dem Fyne-Paket
	sc.setDranIstString()                   // Initialisierung der Anzeige, wer dran ist.
	return sc
}

// Gibt den aktuellen Spieler zurück, wobei false = Spieler 1 ist dran, true = Spieler 2 ist dran, bedeutet.
func (sc *data) GetAktuellerSpieler() bool {
	return sc.aktuellerSpieler
}

// Setzt die aktuellen Spielstände der zwei Spieler in Strings um.
func (sc *data) setScoreStrings() {
	sc.spieler1String.Set("Score: " + strconv.Itoa(sc.scoreSpieler1)) // Der derzeitige Spielstand des Spielers 1 wird als binding.String gespeichert.
	sc.spieler2String.Set("Score: " + strconv.Itoa(sc.scoreSpieler2)) // Der derzeitige Spielstand des Spielers 2 wird als binding.String gespeichert.
}

// Gibt den Spielstand von Spieler 1 als String im Fyne-Paket zurück.
func (sc *data) Spieler1() binding.String {
	return sc.spieler1String
}

// Gibt den Spielstand von Spieler 2 als String im Fyne-Paket zurück.
func (sc *data) Spieler2() binding.String {
	return sc.spieler2String
}

// Setzt die Variable dranIst auf den aktuellen Spieler.
func (sc *data) setDranIstString() {
	if !sc.aktuellerSpieler {
		sc.dranIst.Set("Spieler / Spielerin 1 ist dran!")
	} else {
		sc.dranIst.Set("Spieler / Spielerin 2 ist dran!")
	}
}

// Gibt einen String zurück, der angibt, wer gerade dran ist.
func (sc *data) WerDranIst() binding.String {
	return sc.dranIst
}

// Erhöht den Spielstand des aufrufenden Spielers um 1 und setzt beide Spielstände.
func (sc *data) PaarGefunden() {
	if sc.aktuellerSpieler { // Der aktuellerSpieler = true --> Spieler 2 hat die Funktion aufgerufen.
		sc.scoreSpieler2++ // Der Spielstand von Spieler 2 wird und 1 erhöht.
	} else { // Der aktuellerSpieler = false --> Spieler 1 hat die Funktion aufgerufen.
		sc.scoreSpieler1++ // Der Spielstand von Spieler 1 wird und 1 erhöht.
	}
	sc.setScoreStrings() // Die neuen Spielstände werden gesetzt.
}

// Der nicht aktuelle Spieler wird zum aktuellen Spieler.
func (sc *data) SpielerWechseln() {
	sc.aktuellerSpieler = !sc.aktuellerSpieler // Der jeweils andere Spieler ist mit dem Spielen dran.
	sc.setDranIstString()                      // Aktualisiert die Variable dranIst
}
