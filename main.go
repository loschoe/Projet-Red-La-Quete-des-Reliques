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
		color.HiGreen("|☠️  Boire un poison ([U]        |")
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
			c1.Poisonbottle()
		case "M":
			librairies.Merchant(c1)
		case "F":
			librairies.Forge(c1)
		case "C":
			librairies.CombatMenu(c1)
		case "Exit":
			color.Red("Fermeture du jeu...")
			return
		default:
			color.Red("Choix non reconnu")
		}
		c1.IsDead()

		if c1.GameOver {
			return
		}
	}
}


// ----------------- LANCEMENT ------------------------------
func main() {
	// Lancer le jeu
	librairies.StartGame()

	// Initialiser l'inventaire de base
	inventory := [10]string{
		"Fairy",
		"Miasme",
		"...",
		"...",
		"...",
		"...",
		"...",
		"...",
		"...",
		"...",
	}

	var player librairies.Character

	// Proposer un pseudo personnalisé
	name := librairies.CharacterCreation()

	if name != "" { // Si le joueur a choisi un pseudo
		player = librairies.InitCharacter(
			name,
			"Hylien",
			1,
			100,
			100,
			inventory,
		)
	} else { // Sinon, personnage de base
		player = librairies.InitCharacter(
			"Link",
			"Hylien",
			1,
			500,
			100,
			inventory,
		)
	}

	playerPtr := &player

	fmt.Printf("Bienvenue, %s ! 👋\n", player.Name)
	librairies.DisplayInfo(playerPtr)

	// Lancer le menu principal
	Menu(playerPtr)

	// Si GameOver est true, lancer EndGame
	if playerPtr.GameOver {
		color.Red("Merci d'avoir joué ! Fermeture du jeu...\n")
		librairies.EndGame()
		return
	}
}
