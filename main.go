// Ce fichier est le noyaux central du projet. Le fichier exécutant ! Celui qui va appeler toutes les fonctions 
// Le paquet main puisque c'est l'exécutable 

package main

import (
	"PROJET_RED/librairies"			// Notre librairie contenant les fichiers de code 	
	"fmt"						// Certains prints en dépendent 
	"github.com/fatih/color"		// Afficher des lignes en couleur dans la console 
)

// ---------------- MENU ----------------
// Petite fonction utilitaire : affiche une ligne du menu
func menuItem(text string, c *color.Color) {
	fmt.Println(
		color.WhiteString("|") +
			c.Sprint(text) +
			color.WhiteString("|"),
	)
}

func Menu(c1 *librairies.Character) {
	for {
		librairies.ClearScreen()
		fmt.Println(color.WhiteString("+-------------------------------+"))
		color.White("|             MENU              |")
		fmt.Println(color.WhiteString("+-------------------------------+"))

		menuItem("📜 Infos personnage [P]        ", color.New(color.FgBlue))
		menuItem("🎒 Inventaire [I]              ", color.New(color.FgBlue))
		menuItem("👕 Équipement [E]              ", color.New(color.FgBlue))
		menuItem("🌟 Potion de soin [S]          ", color.New(color.FgGreen))
		menuItem("🍽️  Manger [R]                  ", color.New(color.FgGreen))
		menuItem("☠️  Boire un poison [U]         ", color.New(color.FgHiGreen))
		menuItem("💶 Magasin [M]                 ", color.New(color.FgHiBlack))
		menuItem("⚒️  Forgeron [F]                ", color.New(color.FgHiBlack))
		menuItem("⚔️  Combat [C]                  ", color.New(color.FgHiRed))
		menuItem("                               ", color.New(color.FgHiRed))
		menuItem("🪦  Quitter le jeu [Exit]       ", color.New(color.FgRed))

		fmt.Println(color.WhiteString("+-------------------------------+"))

		color.Yellow("\nVotre choix ? ")
		var choice string
		fmt.Scanln(&choice)

		switch choice {
		case "P":
			librairies.DisplayInfo(c1)
			librairies.Pause()
		case "I":
			c1.AccessInventory()
			librairies.Pause()
		case "E":
			c1.AccessEquipment()
			librairies.Pause()
		case "S":
			c1.TakePot()
			librairies.Pause()
		case "R":
			c1.TakeMeat()
			librairies.Pause()
		case "U":
			c1.Poisonbottle()
		case "M":
			librairies.ClearScreen()
			librairies.Merchant(c1)
		case "F":
			librairies.ClearScreen()
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

	// Initialiser l'inventaire de base du joueur au début du jeu avec la potion offerte 
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

	equipment := [3]string{
		"...",
		"...",
		"...",
	}

	// 1️⃣ Création complète du personnage (nom + classe + stats)
	player := librairies.CreateCharacter()

	// 2️⃣ Ajout de l’inventaire et de l’équipement de départ
	for i := 0; i < len(inventory); i++ {
		if inventory[i] != "..." && inventory[i] != "" {
			player.AddInventory(inventory[i])
		}
	}
	for i := 0; i < len(equipment); i++ {
		if equipment[i] != "..." && equipment[i] != "" {
			player.AddEquipment(equipment[i])
		}
	}

	playerPtr := &player
	player.ApplyEquipmentBonus()

	fmt.Printf("Bienvenue, %s ! 👋\n", player.Name)
	librairies.DisplayInfo(playerPtr)

	// Lancer le menu principal
	Menu(playerPtr)

	// Si GameOver, lancer EndGame
	if playerPtr.GameOver {
		color.Red("Merci d'avoir joué ! Fermeture du jeu...\n")
		librairies.EndGame()
		return
	}
}