package musikAbspieler

import (
	"fmt"
	"os"
	"time"

	beep "github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

type data struct {
	musik    string        //Pfad zur Musikdatei
	streamer beep.Streamer //Gelesene/gestreamte Musikdatei
}

func New(pfad string) *data {
	var m *data
	m = new(data)
	fmt.Println("Das ist der Pfad:", pfad)
	m.musik = pfad
	return m
}

func (m *data) Spielen() {
	// Auslesen des Musikpfades:
	f, err := os.Open(m.musik)
	// Fehlerabfangen, falls der Musikpfad fehlerhaft ist.
	if err != nil {
		fmt.Println("Fataler Fehler: Musikpfad nicht gefunden")
	}

	// Decodierung der MP3-Datei:
	streamer, format, err := mp3.Decode(f)

	fmt.Println("Format", format)
	// Festlegung der Samplerate:
	fmt.Println(format.SampleRate.N(time.Second / 10))
	if err != nil {
		fmt.Println("Fataler Fehler: Fehler beim Decodieren der Musikdatei")
	}

	// Initialisierung des "Lautsprechers"
	m.streamer = streamer

	fmt.Println("Speaker wird initiliziert")

	fmt.Println("ERFOLG! Speaker wurde initiliziert")

	// Musik wird auf dem initialisierten Lautsprecher abgespielt:
	speaker.Play(m.streamer)
}

func (m *data) Stoppen() {
	// Der "Lautsprecher" wird gesäubert/ gelöscht:
	speaker.Clear()
}
