package main

import ("fmt")

// Tâche 5 : Potion de vie

// Créez une fonction takePot qui permet d’utiliser une potion dans l’inventaire. 
// Vous pouvez l’utiliser dans le menu de « Accéder à l’inventaire ». 
// Lorsque vous utilisez une potion, celle-ci se consomme (supprimée de l’inventaire) et vous regagnez 50 points de vie actuel. 
// Puis affichez les points de vie actuel sur les points de vie max du personnage.
// Les points de vie actuels ne peuvent pas excéder les points de vie maximum

func (personnage *Character) TakePot() {
	for i, item := range personnage.Inventory {
		if item == "Fairy" {
			// Soigne le personnage
			personnage.PV += 50
			if personnage.PV > personnage.Max_PV {
				personnage.PV = personnage.Max_PV
			}
			fmt.Println(personnage.Name, "utilise une Fée ! PV =", personnage.PV, "/", personnage.Max_PV)

			// Retire la potion de l'inventaire
			personnage.Inventory[i] = ""
			return
		}
	}
	fmt.Println("Aucune Potion Fée n'est disponible dans l'inventaire.")
}