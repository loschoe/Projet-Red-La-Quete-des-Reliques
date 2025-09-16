package librairies

import (
	"time"
	"github.com/fatih/color"
)

// -------- POTIONS --------
func (personnage *Character) TakePot() {
	for i, item := range personnage.Inventory {
		if item == "Fairy" {
			personnage.PV += 50
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

// Fonction poison buvable 
func (personnage *Character) Poisonbottle() {
	for i, item := range personnage.Inventory {
		if item == "Miasme" {
			println(personnage.Name, "utilise un miasme ! \n")
			for j := 1; j <= 3; j++ {
				time.Sleep(1 * time.Second)
				personnage.PV -= 15
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

func (personnage *Character) PoisonPot(monster *Monster) {
	for i, item := range personnage.Inventory {
		if item == "Miasme" {
			println(personnage.Name, "utilise un miasme ! \n")
			for j := 1; j <= 3; j++ {
				time.Sleep(1 * time.Second)
				monster.PV -= 15
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
