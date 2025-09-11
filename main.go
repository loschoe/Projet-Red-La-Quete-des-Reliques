package main

import "fmt"

// Tâche 1 : Création du personnage
type Character struct {
	Name      string
	Classe    string
	Level     int
	Max_PV    int
	PV        int
	Inventory [10]string
}

// Tâche 2 : Initialisation du personnage
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

func main() {
	// Inventaire au départ !
	var inventory [10]string
	inventory[0] = "Fairy"
	inventory[1] = "Fairy"
	inventory[2] = "Fairy"

	// Initialisation du personnage C1
	c1 := initCharacter("Link", "Hylien", 1, 500, 100, inventory)

	// Affichage du personnage
	fmt.Printf("Name : %s\nClasse : %s\nLevel : %d\nPV : %d/%d\nInventory : %v\n",
		c1.Name, c1.Classe, c1.Level, c1.PV, c1.Max_PV, c1.Inventory)
}
