package main

import (
	"fmt"
	"time"
)

// PoisonPot : inflige 10 PV de dégâts par seconde pendant 3s (30 PV au total)
func (personnage *Character) PoisonPot() {
	fmt.Println(personnage.Name, "utilise une potion de poison !")

	for i := 1; i <= 3; i++ {
		// Attendre 1 seconde
		time.Sleep(1 * time.Second)

		// Infliger 10 points de dégâts
		personnage.PV -= 10
		if personnage.PV < 0 {
			personnage.PV = 0
		}

		// Afficher l’état
		fmt.Printf("Après %d seconde(s) : %d / %d PV\n", i, personnage.PV, personnage.Max_PV)

		// Si le personnage meurt, on arrête
		if personnage.PV == 0 {
			fmt.Println(personnage.Name, "a succombé à ses blessures !")
			return
		}
	}
	fmt.Println("Le poison n’a plus d’effet")
}

