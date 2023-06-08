package spielbrett

import (
	"app/karten"
	"app/score"
	"fmt"
	"math/rand"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"github.com/faiface/beep/speaker"
)

type spielbrett struct { // Das Spielbrett besteht aus:
	karten []karten.Karte // den Karten
	score  score.Score    // den Spielstand
}

func New(score score.Score) *spielbrett {

	speaker.Init(44100, 4410) // Initialisierung des Abspielers, auf die Resource Lautsprecher wird zugegriffen.
	// 44100 - Rate-Sampel-Format, 4410 - Standardrate für MP3

	var sb *spielbrett // Variable Spielbrett wird anlegt
	sb = new(spielbrett)
	sb.karten = make([]karten.Karte, 12) // // Slice für Karten wird angelegt.
	sb.score = score                     // Übergebener Spielstand wird auf dem Spielbrett gespeichert.

	inhaltDerKarten := beliebigerInhalt() // Variable für den Inhalt der Karten wird angelegt.

	// Karten werden mit Inhalt - hier: Musik - gefüllt.
	for i := 0; i < 12; i++ {
		fmt.Println("", i)
		index := i

		// Die Karte wird neu erzeugt. Es wird ihr der Pfad zur Musikdatei mitgegeben sowie, dass sie ausgewählt wurde.
		k := karten.NewKarte("./AudiofilesMemoric/"+inhaltDerKarten[i],
			func() {
				sb.karteAusgewaehlt(index)
			},
		)

		//Die neu erzeugte Karte wird an die i-te Stelle im Slice gelegt:
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

	// Auslesen des Namens der Datei:
	for _, file := range files {
		fmt.Println(file.Name())
	}

	// Bevor Musik nach Anklicken der Karte abgespielt wurde, haben wir uns Buchstaben anzeigen lassen.
	// buchstaben := [6]string{"A", "B", "C", "D", "E", "F"}

	// Anlegen eines Strings, in welchem der bisherige Inhalt doppelt hinterlegt werden soll.
	var verdoppelung []string

	// Der bisherige Inhalt des Slices wird zweimal in den Slice gelegt:
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

	// Das Slice mit zweifach vorliegenden und zufällig angeordneten Dateien wird zurückgegeben.
	return verdoppelung
}

func (sb *spielbrett) GetBrett() fyne.CanvasObject {

	// Die Karten werden in die Fyne-Logik überführt:
	fyneKarten := make([]fyne.CanvasObject, len(sb.karten))
	for i, v := range sb.karten {
		fyneKarten[i] = fyne.CanvasObject(v)
	}

	// Das Spielbrett mit hier festgelegtem Layout aus dem Fyne-Paket wird zurückgegeben.
	return container.New(layout.NewGridLayout(4), fyneKarten...)
}

func (sb *spielbrett) offeneKarten() int {
	// Anlegen der Variable Anzahl mit dem Initialisierungswert 0:
	anzahl := 0
	// Mit der Range-Schleife wird durch das Karten-Slice gegangen und die Anzahl der offenen Karten - offen = true - wird gezählt.
	for _, v := range sb.karten {
		if v.IstOffen() {
			anzahl++
		}
	}

	// Rückgabe der Anzahl der geöffnenten Karten:
	return anzahl
}

func (sb *spielbrett) alleKartenSchliessen() {
	// Mit der Range-Schleife wird durch das Karten-Slice gegangen und alle geöffenten Karten werden geschlossen, indem die Funktion schliessen auf die Karte angewendet wird.
	for _, v := range sb.karten {
		if v.IstOffen() {
			v.Schliessen()
		}
	}
}

// Die Methode "karteAusgewaehlt" wird aufgerufen, wenn eine neue Karte durch einen Spieler ausgewählt wird.
func (sb *spielbrett) karteAusgewaehlt(kartennr int) {
	fmt.Println("Karte wurde ausgewählt")
	fmt.Println("Kartennummer ist: ", kartennr)
	// Sobald eine neue Karte ausgewählt, d.h. angeklickt wird, wird das Abspielen jeglicher Musik gestoppt.
	sb.musikStoppen()

	// Die aufgerufene Karte wird in der Variable ausgewählte Karte gespeichert.
	ausgewaehlteKarte := sb.karten[kartennr]

	//	ausgewaehlteKarte.Verschwinden()

	// Die Methode offeneKarten gibt die Anzahl der offenen Karten zurück, diese Anzahl wird in der Variable anzahlOffen gespeichert.
	anzahlOffen := sb.offeneKarten()

	// Überprüfen, ob die Anzahl der geöffenten Karten kleiner 2 ist. In diesem Fall kann kein Paar vorliegen, da entweder bisher keine oder bisher 1 Karte geöffnet ist.
	// Die angeklickte Karte wird geöffnet.
	if anzahlOffen < 2 {
		ausgewaehlteKarte.Oeffnen()
		// In allen anderen Fällen, also wenn 2 Karten bereits geöffnet wurden, wird überprüft, ob die beiden Karten gleich sind:
	} else {
		if sb.offeneKartenGleich() { // Beide Karten sind gleich:
			sb.score.PaarGefunden()       // Der Zählerstand des aufrufenden Spielers wird um 1 erhöht.
			sb.offeneKartenVerschwinden() // Beide angeklickten Karte verschwinden, d.h. sie werden als "transparent" angezeigt.
		} else { // Beide Karten sind ungleich:
			sb.score.SpielerWechseln() // Es findet ein Spielerwechsel statt.
		}
		sb.alleKartenSchliessen() // Alle geöffneten Karten werden geschlossen.

	}

}

// Die Methode testet, ob alle - hier zwei - geöffneten Karten gleichen Inhalt haben.
func (sb *spielbrett) offeneKartenGleich() bool {
	// Anlegen der Variable "ersterInhalt", in der später der Inhalt der ersten geöffneten Karte gespeichert wird.
	var ersterInhalt string

	// Mit einer Range-Schleife wird durch das Karten-Slice gegangen:
	for _, v := range sb.karten {
		// Überprüfen, ob die Karte geöffent wurde:
		if v.IstOffen() { // Die Karte wurde geöffnet:
			if ersterInhalt == "" { // Wenn "ersterInhalt" noch leer ist, dann wurde die erste geöffnete Karte gefunden.
				ersterInhalt = v.Inhalt() // Der Inhalt der Karte wird in "ersterInhalt" gespeichert.
			} else if ersterInhalt != v.Inhalt() { // Wenn "ersterInhalt" nicht leer ist, d.h. es wurde bereits der Inhalt einer Karte hinterlegt, UND die Inhalte der beiden Karten - Inhalt der ersten Karte in "ersterInhalt" sowie Inhalt der aktuellen Karte - sind nicht gleich
				return false // Rückgabe von false, d.h. beide Karten haben unterschiedliche Inhalte
			}

		}
	}

	return true // Die Inhalte der beiden Karten ist gleich.

}

// Die Methode lässt alle geöffneten Karten verschwinden.
func (sb *spielbrett) offeneKartenVerschwinden() {
	// Mit einer Range-Schleife wird durch alle Karten gegangen.
	for _, v := range sb.karten {
		if v.IstOffen() { // Wenn die Karte geöffnet ist,
			v.Verschwinden() // dann soll die Karte verschwinden.
		}
	}
}

// Die Methode stoppt die Musik von allen Karten.
func (sb *spielbrett) musikStoppen() {
	// Mit einer Range-Schleife wird durch alle Karten gegangen:
	for _, v := range sb.karten {
		v.MusikStoppen()
	}
}
