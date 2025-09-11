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

// Tâche 3 : Affichage des informations du personnage 
func displayInfo(c Character) { 
	fmt.Printf("Name : %s\nClasse : %s\nLevel : %d\nPV : %d/%d\nInventory : %v\n",
		c.Name, c.Classe, c.Level, c.PV, c.Max_PV, c.Inventory)
}

// Tâche 4 : Affichage du contenu de l'inventaire du personnage 
func accessInventory(inventory []string) {
    fmt.Println("Inventaire du personnage :")
    if len(inventory) == 0 {
        fmt.Println("L'inventaire est vide.")
        return
    }
	for i, item := range inventory {
        fmt.Printf("%d. %s\n", i+1, item)
    }
}

// Tâche 5 : Utilisation de la potion de vie
func (personnage *Character) TakePot() {
	for i, item := range personnage.Inventory {
		if item == "Fairy" { // Potion de soin 
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


func main() {
	// Inventaire au départ !
	var inventory [10]string
	inventory[0] = "Fairy"
	inventory[1] = "Fairy"
	inventory[2] = "Fairy"

	// Initialisation du personnage C1
	c1 := initCharacter("Link", "Hylien", 1, 500, 100, inventory)

	// Parcourir l'inventaire et remplacer les cases vides par "..."
	for i, item := range c1.Inventory {
		if item == "" {
			c1.Inventory[i] = "..."
		}
	}

	// Afficher les informations du personnage 
	displayInfo(c1)

}
