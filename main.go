// Ce fichier est le noyaux central du projet. Le fichier exécutant ! Celui qui va appeler toutes les fonctions 
// Le paquet main puisque c'est l'exécutable 

package main

import (
	"PROJET_RED/librairies"			// Notre librairie contenant les fichiers de code 	
	"fmt"							// Certains prints en dépendent 
	"github.com/fatih/color"		// Afficher des lignes en couleur dans la console 
)

// ---------------- MENU ----------------
func Menu(c1 *librairies.Character) {
	for {
		fmt.Println("+-------------------------------+")
		color.Cyan("|             MENU              |")
		fmt.Println("+-------------------------------+")
		color.Blue("|📜 Infos personnage [P]        |")		// Diplay info 
		color.Blue("|🎒 Inventaire [I]              |")		// AccessInventory
		color.Blue("|👕 Équipement [E]              |")      // AccessEquipment
		color.Green("|🌟 Potion de soin [S]          |")	// TakePot	
		color.Green("|🍽️  Manger [R]                  |")   // TakeMeat
		color.HiGreen("|☠️  Boire un poison [U]         |")	// PoisonBottle 
		color.HiBlack("|💶 Magasin [M]                 |")	// Merchant 
		color.HiBlack("|⚒️  Forgeron [F]                |")	// Forge 
		color.HiRed("|⚔️  Combat [C]                  |")	 // Combat Menu 	
		color.HiRed("|                               |")
		color.Red("|🪦  Quitter le jeu [Exit]       |")		// Exit 
		fmt.Println("+-------------------------------+")

		color.Yellow("\nVotre choix ? ")
		var choice string
		fmt.Scanln(&choice)

		switch choice {
		case "P":
			librairies.DisplayInfo(c1)
		case "I":
			c1.AccessInventory()
		case "E":
			c1.AccessEquipment()
		case "S":
			c1.TakePot()
		case "R":
			c1.TakeMeat()	
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
// ----------------- LANCEMENT ------------------------------
func main() {
	// Lancer le jeu
	librairies.StartGame()

	// Initialiser l'inventaire de base du joueur au début du jeu avec la potion offerte 
	inventory := [10]string{
		"Fairy",
		"Master Sword",
		"Tissu royal",
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
