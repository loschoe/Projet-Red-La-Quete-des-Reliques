// Ce fichier contient toutes les fonctions nécéssaires au fonctionnement des items qui sont présent dans le jeu. 
// Le paquet de la librairie où sont stockées les fonctions. 

package librairies

import (
	"github.com/fatih/color" // Couleurs 
)

//=================== EPEE ==================================
// Nombre d'utilisation de l'épée
var MasterSwordUses = 100

func (personnage *Character) UseMasterSword(monster *Monster) {
	for i, item := range personnage.Inventory {
		if item == "Master Sword" {
			println(personnage.Name, "utilise l'épée purificatrice !\n")

			// Inflige 50 points de dégâts 
			monster.PV -= 50
			if monster.PV < 0 {
				monster.PV = 0
			}
			color.Red("%s inflige 50 points de dégâts ! PV restant du monstre : %d / %d\n",
				personnage.Name, monster.PV, monster.Max_PV)

			// Réduire le compteur d'usage 
			if MasterSwordUses > 0 {
				MasterSwordUses--
				color.HiBlack("Usages restants de la Master Sword : %d\n", MasterSwordUses)
			} else {
				color.HiBlack("La Master Sword est usée !\n")
				// Retirer l'épée si usages = 0
				personnage.RemoveItemAt(i)
			}

			return
		}
	}
	color.Red("Vous n'avez pas l'épée dans votre inventaire.\n")
}

//=================== ARC ==========================
var BowUses = 4 // Compteur d'usage 

func (personnage *Character) UseBow(monster *Monster) {
	// Vérifie si le personnage a un arc
	arcIndex := -1
	for i, item := range personnage.Inventory {
		if item == "Bow" {
			arcIndex = i
			break
		}
	}

	if arcIndex == -1 {
		color.Red("Vous n'avez pas l'arc dans votre inventaire.\n")
		return
	}

	// Vérifie si le personnage a une flèche
	arrowIndex := -1
	for i, item := range personnage.Inventory {
		if item == "Arrow" {
			arrowIndex = i
			break
		}
	}

	if arrowIndex == -1 {
		color.Red("Vous n'avez plus de flèches.\n")
		return
	}

	println(personnage.Name, "utilise un arc !\n")

	// Inflige 100 points de dégâts
	monster.PV -= 100
	if monster.PV < 0 {
		monster.PV = 0
	}
	color.Red("%s inflige 100 points de dégâts ! PV restant du monstre : %d / %d\n",
		personnage.Name, monster.PV, monster.Max_PV)

	// Réduire le compteur d'usage de l'arc
	if BowUses > 0 {
		BowUses--
		color.HiBlack("Usages restants de l'arc : %d\n", BowUses)
	} else {
		color.HiBlack("L'arc est usé !\n")
		// Retirer l'arc si usages = 0
		personnage.RemoveItemAt(arcIndex)
	}

	// Retirer la flèche utilisée
	personnage.RemoveItemAt(arrowIndex)
}
