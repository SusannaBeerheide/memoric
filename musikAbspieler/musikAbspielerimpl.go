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
	m.musik = pfad
	return m
}

func (m *data) Spielen() {
	f, err := os.Open(m.musik)
	if err != nil {
		fmt.Println("Fataler Fehler: Musikpfad nicht gefunden")
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		fmt.Println("Fataler Fehler: Fehler beim Decodieren der Musikdatei")
	}

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	speaker.Play(streamer)
}

func (m *data) Stoppen() {

}
