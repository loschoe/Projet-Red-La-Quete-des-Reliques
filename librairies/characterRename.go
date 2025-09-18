// Ce fichier contient toutes les fonctions nécéssaires pour rennomer le personnage  
// Le paquet de la librairie où sont stockées les fonctions 
package librairies

import (
	"fmt" // Pour les affichages 
	"github.com/fatih/color"
)

// Focntion qui vérifie si la chaîne de caractères contient uniquement des lettres
func IsAlpha(s string) bool {
	for i := 0; i < len(s); i++ {
		if !((s[i] >= 'A' && s[i] <= 'Z') || (s[i] >= 'a' && s[i] <= 'z')) {
			return false
		}
	}
	return true
}

// Fonction qui met la première lettre en majuscule et le reste en minuscule
func FormatName(s string) string {
	if len(s) == 0 {
		return s
	}
	b := []byte(s)
	if b[0] >= 'a' && b[0] <= 'z' {
		b[0] -= 32
	}
	for i := 1; i < len(b); i++ {
		if b[i] >= 'A' && b[i] <= 'Z' {
			b[i] += 32
		}
	}
	return string(b)
}

// Fonction principale de la création d'un pseudo personnalisé 
func CharacterCreation() string {
	var input, confirm string
	for {
		color.Cyan("Voulez-vous un pseudo personnalisé ? (y/n) : ")
		fmt.Scanln(&confirm)

		if confirm == "y" || confirm == "oui" {
			// Saisie du pseudo
			for {
				color.Cyan("Entrez le nom de votre personnage : ")
				fmt.Scanln(&input)

				if !IsAlpha(input) {
					color.Red("Erreur : le nom ne doit contenir que des lettres.")
					continue
				}

				name := FormatName(input)
				fmt.Printf("Nom proposé : %s\n", name)

				color.Cyan("Voulez-vous garder ce nom ? (y/n) : ")
				fmt.Scanln(&confirm)

				if confirm == "y" || confirm == "oui" {
					return name
				} else if confirm == "n" || confirm == "non" {
					color.Red("Recommencez la saisie du nom.")
					continue
				} else {
					color.Red("Réponse invalide, veuillez répondre par 'y' ou 'n'.")
				}
			}
		} else if confirm == "n" || confirm == "non" {
			// Le joueur ne veut pas de pseudo personnalisé
			return ""
		} else {
			color.Red("Réponse invalide, veuillez répondre par 'y' ou 'n'.")
		}
	}
}
