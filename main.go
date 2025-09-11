package main

import (
	"fmt"
	"github.com/fatih/color"
	"time"
)

// Fonction de démarrage
func startGame() {
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

	// Attendre que l’utilisateur appuie sur Entrée
	fmt.Scanln()
}

// Définition du personnage
type Character struct {
	Name      string
	Classe    string
	Level     int
	Max_PV    int
	PV        int
	Inventory [10]string
}

// Initialisation du personnage
func initCharacter(name string, classe string, level int, max_pv int, pv int, inventory [10]string) Character {
	if pv > max_pv {
		pv = max_pv
	}
	return Character{
		Name:      name,
		Classe:    classe,
		Level:     level,
		Max_PV:    max_pv,
		PV:        pv,
		Inventory: inventory,
	}
}

// Affichage des informations
func displayInfo(c Character) {
	fmt.Printf("\nName : %s\nClasse : %s\nLevel : %d\nPV : %d/%d\nInventory : %v\n",
		c.Name, c.Classe, c.Level, c.PV, c.Max_PV, c.Inventory)
}

// Affichage de l'inventaire
func accessInventory(inventory [10]string) {
	fmt.Println("\nInventaire du personnage :")
	empty := true
	for i, item := range inventory {
		if item != "..." && item != "" {
			fmt.Printf("%d. %s\n", i+1, item)
			empty = false
		}
	}
	if empty {
		fmt.Println("L'inventaire est vide.")
	}
}

// Utilisation d'une potion
func (personnage *Character) TakePot() {
	for i, item := range personnage.Inventory {
		if item == "Fairy" {
			personnage.PV += 50
			if personnage.PV > personnage.Max_PV {
				personnage.PV = personnage.Max_PV
			}
			fmt.Println(personnage.Name, "utilise une Fée ! PV =", personnage.PV, "/", personnage.Max_PV)
			personnage.Inventory[i] = "..."
			return
		}
	}
	fmt.Println("Aucune Potion Fée n'est disponible dans l'inventaire.")
}

// Tâche 9 : inflige 10 PV de dégâts par seconde pendant 3s (30 PV au total)
func (personnage *Character) PoisonPot() {
	for i, item := range personnage.Inventory {
		if item == "Miasme" {
			fmt.Println(personnage.Name, "utilise un miasme !")

			for j := 1; j <= 3; j++ {
				// Attendre 1 seconde
				time.Sleep(1 * time.Second)

				// Infliger 15 points de dégâts
				personnage.PV -= 15
				if personnage.PV < 0 {
					personnage.PV = 0
				}

				// Afficher l’état
				fmt.Printf("Après %d seconde(s) : %d / %d PV\n", j, personnage.PV, personnage.Max_PV)

				// Si le personnage meurt, on arrête
				if personnage.PV == 0 {
					fmt.Println(personnage.Name, "a succombé à ses blessures !")
					personnage.RemoveItemAt(i) // 🔥 retirer le miasme
					return
				}
			}

			fmt.Println("Le miasme n’a plus d’effet")
			personnage.RemoveItemAt(i) // 🔥 retirer le miasme
			return
		}
	}

	fmt.Println("Aucun Miasme n'est disponible dans l'inventaire.")
}

func (personnage *Character) RemoveItemAt(index int) {
	for j := index; j < len(personnage.Inventory)-1; j++ {
		personnage.Inventory[j] = personnage.Inventory[j+1]
	}
	personnage.Inventory[len(personnage.Inventory)-1] = "" // vide la dernière case
}


// Fonction menu
func menu(c1 Character) {
	for {
		color.Cyan("\nMENU")
		color.Blue("- Informations personnage [P]")
		color.Blue("- Accéder à l’inventaire [I]")
		color.Green("- Utiliser une potion [U]")
		color.Green("- Magasin [M]")
		color.Red("\n - Quitter le jeu [Exit]")

		var choice string
		color.Yellow("\nVers quel menu souhaitez-vous aller ? ")
		fmt.Scanln(&choice)

		switch choice {
		case "P":
			displayInfo(c1)
		case "I":
			accessInventory(c1.Inventory)
		case "U":
			c1.TakePot()
		case "M":
			c1.PoisonPot()
		case "Exit":
			color.Red("Fermeture du jeu...")
			return
		default:
			color.Red("Choix non reconnu")
		}
	}
}

// Fonction main
func main() {
	// Étape 1 : démarrage avec ascii art + Entrée
	startGame()

	// Inventaire initial
	var inventory [10]string
	inventory[0] = "Fairy"
	inventory[1] = "Fairy"
	inventory[2] = "Fairy"
	inventory[3] = "Miasme"

	c1 := initCharacter("Link", "Hylien", 1, 500, 100, inventory)

<<<<<<< HEAD
    // Tâche 6 : Menu 
    for {
        color.Cyan("\nMENU")
        color.Blue("Informations personnage [P]")
        color.Blue("Accéder à l’inventaire [I]")
        color.Blue("Utiliser une potion de soin [S]")
        color.color.HiGreen("Utiliser une potion de poison [O]")
        color.Red("Quitter le jeu [Exit]")
=======
	// Remplacer les cases vides par "..."
	for i, item := range c1.Inventory {
		if item == "" {
			c1.Inventory[i] = "..."
		}
	}
>>>>>>> 7df976a00bead96f40b038bdd21160441019a65e

	// Étape 2 : lancement du menu
	menu(c1)
}

