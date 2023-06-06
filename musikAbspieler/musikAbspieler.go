package musikAbspieler

//import

// Vor.: Übergabe des Pfad der Musikdatei.
// Erg.: Eine neue Instanz des ADT Musikabspieler ist gegeben.
// func New(pfad string) *data // *data erfüllt das Interface Schloss

type MusikAbspieler interface {
	// Vor.: keine
	// Eff.: Falls der übergebene Pfad zu einer Musikdatei im MP3-Format führt, wird diese abgespielt.
	Spielen()

	// Vor.: keine
	// Eff.: Das Abspielen der Musik ist gestoppt.
	Stoppen()
}
