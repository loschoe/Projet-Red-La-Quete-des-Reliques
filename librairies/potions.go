// Ce fichier contient toutes les fonctions nécéssaires aux potions et à leur fonctionnement 
// Le paquet de la librairie où sont stockées les fonctions 

package librairies

import (
	"time" 						//le temps d'empoisonnement 
	"github.com/fatih/color"	//afficher en couleurs 
)

// -------- POTION DE SOIN --------
func (personnage *Character) TakePot() {
	for i, item := range personnage.Inventory {
		if item == "Fairy" {
			personnage.PV += 55
			if personnage.PV > personnage.Max_PV {
				personnage.PV = personnage.Max_PV
			}
			println(personnage.Name, "utilise une Fée ! PV =", personnage.PV, "/", personnage.Max_PV)
			personnage.Inventory[i] = "..."
			return
		}
	}
	color.Red("Aucune Fée disponible.")
}

// -------- Venaison divine --------
func (personnage *Character) TakeMeat() {
	for i, item := range personnage.Inventory {
		if item == "Divine Venison" {
			personnage.PV += 20
			if personnage.PV > personnage.Max_PV {
				personnage.PV = personnage.Max_PV
			}
			println(personnage.Name, "déguste une pièce de Boeuf d'exception ! PV =", personnage.PV, "/", personnage.Max_PV)
			personnage.Inventory[i] = "..."
			return
		}
	}
	color.Red("Aucune Fée disponible.")
}

// -------- POTION DE POISON (buvable) --------
func (personnage *Character) Poisonbottle() {
	for i, item := range personnage.Inventory {
		if item == "Miasme" {
			println(personnage.Name, "utilise un miasme ! \n")
			for j := 1; j <= 3; j++ {
				time.Sleep(1 * time.Second)
				personnage.PV -= 15				// Les dégats sont transmis au joueur 
				if personnage.PV < 0 {
					personnage.PV = 0
				}
				color.Red("Après %d seconde(s) : %d / %d PV\n", j, personnage.PV, personnage.Max_PV)
			}
			color.HiBlack("Le miasme n’a plus d’effet \n")
			personnage.RemoveItemAt(i)
			return
		}
	}
	color.Red("Aucun Miasme disponible.\n")
}

// -------- POTION DE POISON (Lançable) --------
func (personnage *Character) PoisonPot(monster *Monster) {
	for i, item := range personnage.Inventory {
		if item == "Miasme" {
			println(personnage.Name, "utilise un miasme ! \n")
			for j := 1; j <= 3; j++ {
				time.Sleep(1 * time.Second)
				monster.PV -= 15			// Les dégâts sont transmis au monstre 
				if monster.PV < 0 {
					monster.PV = 0
				}
				color.Red("Après %d seconde(s) : %d / %d PV\n", j, monster.PV, monster.Max_PV)
			}
			color.HiBlack("Le miasme n’a plus d’effet \n")
			personnage.RemoveItemAt(i)
			return
		}
	}
	color.Red("Aucun Miasme disponible.\n")
}
