package main

import "fmt"

// accessInventory affiche le contenu de l'inventaire du personnage (Tâche 4)
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

