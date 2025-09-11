package main

import (
    "fmt"
    "github.com/fatih/color"
)

// Définition du personnage
type Character struct {
    Name      string
    Classe    string
    Level     int
    Max_PV    int
    PV        int
    Inventory [10]string
}

// Initialisation du personnage
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

// Affichage des informations
func displayInfo(c Character) {
    fmt.Printf("\nName : %s\nClasse : %s\nLevel : %d\nPV : %d/%d\nInventory : %v\n",
        c.Name, c.Classe, c.Level, c.PV, c.Max_PV, c.Inventory)
}

// Affichage de l'inventaire
func accessInventory(inventory [10]string) {
    fmt.Println("\nInventaire du personnage :")
    empty := true
    for i, item := range inventory {
        if item != "..." && item != "" {
            fmt.Printf("%d. %s\n", i+1, item)
            empty = false
        }
    }
    if empty {
        fmt.Println("L'inventaire est vide.")
    }
}

// Utilisation d'une potion
func (personnage *Character) TakePot() {
    for i, item := range personnage.Inventory {
        if item == "Fairy" {
            personnage.PV += 50
            if personnage.PV > personnage.Max_PV {
                personnage.PV = personnage.Max_PV
            }
            fmt.Println(personnage.Name, "utilise une Fée ! PV =", personnage.PV, "/", personnage.Max_PV)
            personnage.Inventory[i] = "..."
            return
        }
    }
    fmt.Println("Aucune Potion Fée n'est disponible dans l'inventaire.")
}

// Fonction main
func main() {
    // Inventaire initial
    var inventory [10]string
    inventory[0] = "Fairy"
    inventory[1] = "Fairy"
    inventory[2] = "Fairy"

    c1 := initCharacter("Link", "Hylien", 1, 500, 100, inventory)

    // Remplacer les cases vides par "..."
    for i, item := range c1.Inventory {
        if item == "" {
            c1.Inventory[i] = "..."
        }
    }

    // Tâche 6 : Menu 
    for {
        color.Cyan("\nMENU")
        color.Blue("Informations personnage [P]")
        color.Blue("Accéder à l’inventaire [I]")
        color.Blue("Utiliser une potion de soin [S]")
        color.color.HiGreen("Utiliser une potion de poison [O]")
        color.Red("Quitter le jeu [Exit]")

        var choice string
        color.Yellow("\nVers quel menu souhaitez-vous aller ? ")
        fmt.Scanln(&choice)

        switch choice {
        case "P":
            displayInfo(c1)
        case "I":
            accessInventory(c1.Inventory)
        case "U":
            c1.TakePot()
        case "Exit":
            color.Red("Fermeture du jeu...")
            return
        default:
            color.Red("Choix non reconnu")
        }
    }
}
