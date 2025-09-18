// Ce fichier contient toutes les fonctions nécéssaires au démarrage et à la cloture esthétique du jeu.
// Le paquet de la librairie où sont stockées les fonctions

package librairies

import (
	"fmt" // Prints

	"github.com/fatih/color" // Couleurs dans la console
)

// ----------- Menu DEMARRER ---------------------
func StartGame() {
	asciiArt := `
                           |>>>
            |>>>       _  _|_  _         |>>>
            |         |;|_|;|_|;|        |
        _  _|_  _     \         /    _  _|_  _
       |;|_|;|_|;|     \       /    |;|_|;|_|;|
       \ ..      /     ||     |     \         /
	\ .     /      ||     |      \       /
	||:	|_   _ ||_  _ |  _   _||:    |
	||:	|||_|;|_|;|_|;|_|;|_|;||:    |
	||:	||                    ||:    |
	||:	||                    ||:    |
	||:	||      _______       ||:    |
	||:	||     /+++++++\      ||:    |
	||:	||     |+++++++|      ||:    |
     __	||:	||     |+++++++|     _||_    |
___--	'--~~____|     |+++++__|----~    ~---,
		 ~---__|,--~'                  ~~---
`
	introText := `
      ______ _              ____           _   _
     /__    \ |__  ___     / ___\___ _ ___| |_| | ___
       / /\/  _  \/ _ \   / /   /   ' / __| __| |/ _ \
      / /  | | | |  __/  / /___|  (_| \__ \ |_| |  __/
      \/   |_| |_|\___|  \_____/\___,_|___/\__|_|\___|

	  Appuyer sur Entrée pour commencer !!
	`

	color.Cyan("%s\n", asciiArt)
	color.Red("%s\n", introText)
	fmt.Scanln()
}

// ----------- Menu END Game --------------------- (On remarquera que le chateau est ouvert !)
func EndGame() {
	asciiArtt := `
                           |>>>
            |>>>       _  _|_  _         |>>>
            |         |;|_|;|_|;|        |
        _  _|_  _     \         /    _  _|_  _
       |;|_|;|_|;|     \       /    |;|_|;|_|;|
       \ ..      /     ||     |     \         /
	\ .     /      ||     |      \       /
	||:	|_   _ ||_  _ |  _   _||:    |
	||:	|||_|;|_|;|_|;|_|;|_|;||:    |
	||:	||                    ||:    |
	||:	||                    ||:    |
	||:	||      _______       ||:    |
	||:	||     /       \      ||:    |
	||:	||     |       |      ||:    |
     __	||:	||     |       |     _||_    |
___--	'--~~____|     |     __|----~    ~---,
		 ~---__|,--~'                  HF---
`
	conclText := `
      ______ _              ____           _   _
     /__    \ |__  ___     / ___\___ _ ___| |_| | ___
       / /\/  _  \/ _ \   / /   /   ' / __| __| |/ _ \
      / /  | | | |  __/  / /___|  (_| \__ \ |_| |  __/
      \/   |_| |_|\___|  \_____/\___,_|___/\__|_|\___|

	  Bravo vous avez terminer le jeu !!
	`
	ClearScreen()
	color.Yellow("%s\n", asciiArtt)
	color.Red("%s\n", conclText)
	fmt.Scanln()
}
