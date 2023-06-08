package main

import (
	"app/musikAbspieler"
	"fmt"
	"time"

	"github.com/faiface/beep/speaker"
)

func main() {
	speaker.Init(44100, 4410)
	musik := musikAbspieler.New("../AudiofilesMemoric/Vogelhaus.mp3")
	time.Sleep(100)
	musik.Spielen()

	fmt.Println("Jetzt spielt die Musik")

}
