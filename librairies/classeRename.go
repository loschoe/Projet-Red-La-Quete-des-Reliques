// Ce fichier contient toutes les fonctions nécéssaires pour personnaliser son personnage et changer sa classe.
// Le paquet de la librairie où sont stockées les fonctions 
package librairies

import (
	"fmt"
	"github.com/fatih/color"
)

// ---------------- CHOIX DE CLASSE ----------------

// Permet au joueur de choisir sa classe
func ChooseClass() string {
	var input, confirm string
	classes := []string{"Hylien", "Zora", "Goron"}

	for {
		color.Cyan("Choisissez la classe de votre personnage : Hylien, Zora ou Goron.")
		fmt.Scanln(&input)

		input = FormatName(input) // Utilise la fonction déjà définie dans characterRename.go

		// Vérifie que la classe est valide
		valid := false
		for _, c := range classes {
			if input == c {
				valid = true
				break
			}
		}

		if !valid {
			color.Red("Classe invalide ! Veuillez choisir parmi : Hylien, Zora ou Goron.")
			continue
		}

		color.Cyan("Classe choisie : %s\n", input)
		color.Cyan("Voulez-vous confirmer cette classe ? (y/n) : ")
		fmt.Scanln(&confirm)

		if len(confirm) > 0 {
			c := confirm[0]
			if c >= 'A' && c <= 'Z' {
				c += 32
			}
			confirm = string(c)
		}

		if confirm == "y" || confirm == "o" { // 'o' pour "oui"
			return input
		} else if confirm == "n" {
			color.Red("Recommencez la sélection de la classe.")
			continue
		} else {
			color.Red("Réponse invalide, veuillez répondre par 'y' ou 'n'.")
		}
	}
}

// ---------------- CREATION DE PERSONNAGE ----------------

func CreateCharacter() Character {
	// Choix du pseudo
	name := CharacterCreation()
	if name == "" {
		name = "Link" // Nom par défaut si pas de pseudo
	}

	// Choix de la classe
	classe := ChooseClass()

	// Initialisation des caractéristiques selon la classe
	var maxPV, pv, attack, rubis int
	var skills []string
	rubis = 100

	switch classe {
	case "Hylien":
    	maxPV = 200       // PV max pour Hylien
    	pv = 100          // PV actuels de départ (50%)
    	attack = 6
    	skills = []string{"Coup de Poing"}
	case "Zora":
    	maxPV = 250       // PV max pour Zora
    	pv = 150          // PV actuels de départ
    	attack = 10
    	skills = []string{"Aquatique"}
		rubis = 200       // Bonus de rubis pour Zora
	case "Goron":
    	maxPV = 300       // PV max pour Goron
    	pv = 200          // PV actuels de départ
    	attack = 4
    	skills = []string{"Roulade"}
	default:
    	maxPV = 200
    	pv = 100
    	attack = 6
    	skills = []string{"Coup de Poing"}
	}

	// Inventaire et équipement par défaut
	inventory := [10]string{
		"...", 
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
	equipment := [3]string{"...", "...", "..."}

	// Création du personnage
	player := InitCharacter(name, classe, 1, maxPV, pv, inventory, equipment)
	player.Attack = attack
	player.Skills = skills
	player.Rubis = rubis
	player.FireBallUsed = false

	return player
}