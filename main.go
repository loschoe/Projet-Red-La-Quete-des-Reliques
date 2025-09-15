package main

import (
	"PROJET_RED/librairies"
	"fmt"
	"github.com/fatih/color"
)

// ---------------- MENU ----------------
func Menu(c1 *librairies.Character) {
	for {
		fmt.Println("+-------------------------------+")
		color.Cyan("|             MENU              |")
		fmt.Println("+-------------------------------+")
		color.Blue("|👕 Infos personnage [P]        |")
		color.Blue("|🎒 Inventaire [I]              |")
		color.Green("|🌟 Potion de soin [S]          |")
		color.HiGreen("|☠️  Potion de poison [U]        |")
		color.HiBlack("|💶 Magasin [M]                 |")
		color.HiBlack("|⚔️  Forgeron [F]                |")
		color.HiRed("|🛡️  Combat [C]                  |")
		color.HiRed("|                               |")
		color.Red("|❌ Quitter le jeu [Exit]       |")
		fmt.Println("+-------------------------------+")

		color.Yellow("\nVotre choix ? ")
		var choice string
		fmt.Scanln(&choice)

		switch choice {
		case "P":
			librairies.DisplayInfo(c1)
		case "I":
			c1.AccessInventory()
		case "S":
			c1.TakePot()
		case "U":
			c1.PoisonPot()
		case "M":
			librairies.Merchant(c1)
		case "F":
			librairies.Forge(c1)
		case "C":
			librairies.TrainingFight(c1)
		case "Exit":
			color.Red("Fermeture du jeu...")
			return
		default:
			color.Red("Choix non reconnu")
		}
		c1.IsDead()
	}
}


// ----------------- LANCEMENT ------------------------------
func main() {
	// Lancer le jeu
	librairies.StartGame()

	// Initialiser l'inventaire de base
	inventory := [10]string{
		"Fairy",
		"...",
		"...",
		"...",
		"...",
		"...",
		"...",
		"...",
		"...",
		"...",	
	}

	// Créer le personnage
	c := librairies.InitCharacter("Link", "Hylien", 1, 500, 100, inventory)
	c1 := &c

	// Lancer le menu principal
	Menu(c1)
}