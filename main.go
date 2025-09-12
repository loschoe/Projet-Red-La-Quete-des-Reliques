package main

import (
	"fmt"
	"github.com/fatih/color"
	"time"
)

// Fonction de démarrage
func startGame() {
	asciiArt := `
                           |>>>
            |>>>       _  _|_  _         |>>>
            |         |;|_|;|_|;|        |
        _  _|_  _     \         /    _  _|_  _
       |;|_|;|_|;|     \       /    |;|_|;|_|;|
       \ ..      /     ||     |     \         /
	\ .     /      ||     |      \       /
	||:	|_   _ ||_  _ |  _   _||:    |
	||:	|||_|;|_|;|_|;|_|;|_|;||:    |
	||:	||                    ||:    |
	||:	||                    ||:    |
	||:	||      _______       ||:    |
	||:	||     /+++++++\      ||:    |
	||:	||     |+++++++|      ||:    |
     __	||:	||     |+++++++|     _||_    |
___--	'--~~____|     |+++++__|----~    ~---,
		 ~---__|,--~'                  ~~---
`
	introText := `
      ______ _              ____           _   _
     /__    \ |__  ___     / ___\___ _ ___| |_| | ___
       / /\/  _  \/ _ \   / /   /   ' / __| __| |/ _ \
      / /  | | | |  __/  / /___|  (_| \__ \ |_| |  __/
      \/   |_| |_|\___|  \_____/\___,_|___/\__|_|\___|

	  Appuyer sur Entrée pour commencer !!
	`

	color.Cyan("%s\n", asciiArt)
	color.Red("%s\n", introText)

	fmt.Scanln()
}

// Tâche 1 : Définition du personnage 
type Character struct {
	Name               string
	Classe             string
	Level              int
	Max_PV             int
	PV                 int
	Inventory          [10]string
	HasReceivedDiamond bool
	Rubis              int
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
		Rubis:     10,
	}
}

// Tâche 3 : Affichage des informations
func displayInfo(c Character) {
	fmt.Printf("\nName : %s\nClasse : %s\nLevel : %d\nPV : %d/%d\nInventory : %v\n",
		c.Name, c.Classe, c.Level, c.PV, c.Max_PV, c.Inventory)
}

// Afficher l'inventaire
func (personnage *Character) accessInventory() {
    fmt.Println("\nInventaire du personnage :")
    vide := true
    for i, item := range personnage.Inventory {
        if item == "" || item == "..." {
            fmt.Printf("%d. [vide]\n", i+1)
        } else {
            fmt.Printf("%d. %s (utilisable)\n", i+1, item)
            vide = false
        }
    }
    if vide {
        fmt.Println("L'inventaire de votre personnage est vide.")
    }
    fmt.Println("Rubis disponibles :", personnage.Rubis)
    fmt.Println()
}

// Tâche 5 : Utilisation d'une potion de soin (renommée Fairy)
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
	fmt.Println("Aucune Fée n'est disponible dans l'inventaire.")
}

// Ajouter un item dans la première case libre ("" ou "...")
func (personnage *Character) AddInventory(item string) {
    for i := 0; i < len(personnage.Inventory); i++ {
        if personnage.Inventory[i] == "" || personnage.Inventory[i] == "..." {
            personnage.Inventory[i] = item
            fmt.Println(item, "a été ajouté à l'inventaire avec succès !")
            return
        }
    }
    fmt.Println("Inventaire plein ! Impossible d'ajouter l'item", item)
}

// Tâche 7 : Supprimer le premier exemplaire d'un item précis
func (personnage *Character) RemoveItem(item string) {
	for idx, i := range personnage.Inventory {
		if i == item {
			personnage.Inventory[idx] = ""
			fmt.Println(item, "a été retiré de l'inventaire.")
			return
		}
	}
	fmt.Println("Aucun(e)", item, "n'a été trouvé dans l'inventaire.")
}

// Compter combien d'exemplaires d'un item précis
func (personnage *Character) CountItem(item string) int {
	count := 0
	for _, i := range personnage.Inventory {
		if i == item {
			count++
		}
	}
	return count
}

// La fonction du marchand

func Merchant(personnage *Character) {
	fmt.Println("\nBienvenue dans ma boutique !")

	type ShopItem struct {
		Name       string
		PriceRubis int
		PriceDiam  int
		EffectPV   int
	}

	shopItems := []ShopItem{
		{"5X Arrow", 10, 0, 0},
		{"5X Arrow", 10, 0, 0},
		{"Master Sword", 0, 3, 0},
		{"PoisonPot", 25, 0, 0},
		{"Divine Venison", 50, 0, 25},
	}

	if !personnage.HasReceivedDiamond {
		personnage.AddInventory("Diamant")
		personnage.HasReceivedDiamond = true
		fmt.Println("Vous avez reçu un Diamant en cadeau !")
	}

	for len(shopItems) > 0 {
		fmt.Println("\nItems disponibles :")
		for i, item := range shopItems {
			fmt.Printf("%d. %s", i+1, item.Name)
			if item.PriceRubis > 0 {
				fmt.Printf(" - %d Rubis", item.PriceRubis)
			}
			if item.PriceDiam > 0 {
				fmt.Printf(" - %d Diamants", item.PriceDiam)
			}
			fmt.Println()
		}
		fmt.Println("0. Quitter")
		fmt.Println("Pour débloquer le Bouclier d'Hylia, vous devez terminer 5 entraînements.")
		fmt.Println("Rubis disponibles :", personnage.Rubis)

		var choix int
		fmt.Print("Votre choix : ")
		fmt.Scan(&choix)

		if choix == 0 {
			fmt.Println("Au revoir !")
			return
		}

		if choix < 1 || choix > len(shopItems) {
			fmt.Println("Choix invalide.")
			continue
		}

		selectedItem := shopItems[choix-1]

		if selectedItem.PriceRubis > personnage.Rubis {
			fmt.Println("Vous n'avez pas assez de Rubis pour acheter cet item !")
			continue
		}
		if selectedItem.PriceDiam > personnage.CountItem("Diamant") {
			fmt.Println("Vous n'avez pas assez de Diamants pour obtenir cet item !")
			continue
		}

		personnage.Rubis -= selectedItem.PriceRubis
		for i := 0; i < selectedItem.PriceDiam; i++ {
			personnage.RemoveItem("Diamant")
		}

		personnage.AddInventory(selectedItem.Name)

		if selectedItem.EffectPV > 0 {
			personnage.PV += selectedItem.EffectPV
			if personnage.PV > personnage.Max_PV {
				personnage.PV = personnage.Max_PV
			}
			fmt.Printf("%s vous rend %d PV ! PV actuel : %d/%d\n", selectedItem.Name, selectedItem.EffectPV, personnage.PV, personnage.Max_PV)
		}

		fmt.Println(selectedItem.Name, "a été ajouté à votre inventaire !")

		shopItems = append(shopItems[:choix-1], shopItems[choix:]...)
	}

	fmt.Println("Le marchand n'a plus d'items à vendre. Retour au menu principal.")
}

// Tâche 9 : Utilisation d'une potion de poison (renommée miasme)
func (personnage *Character) PoisonPot() {
	for i, item := range personnage.Inventory {
		if item == "Miasme" {
			fmt.Println(personnage.Name, "utilise un miasme !")

			for j := 1; j <= 3; j++ {
				// Attendre 1 seconde
				time.Sleep(1 * time.Second)

				// Infliger 15 points de dégâts
				personnage.PV -= 15
				if personnage.PV < 0 {
					personnage.PV = 0
				}

				// Afficher l’état
				fmt.Printf("Après %d seconde(s) : %d / %d PV\n", j, personnage.PV, personnage.Max_PV)
			}

			fmt.Println("Le miasme n’a plus d’effet")
			personnage.RemoveItemAt(i)
			return
		}
	}

	fmt.Println("Aucun Miasme n'est disponible dans l'inventaire.")
}

func (personnage *Character) RemoveItemAt(index int) {
	for j := index; j < len(personnage.Inventory)-1; j++ {
		personnage.Inventory[j] = personnage.Inventory[j+1]
	}
	personnage.Inventory[len(personnage.Inventory)-1] = ""
}

// Fonction menu
func menu(c1 *Character){
	for {
	fmt.Println("+-------------------------------+")
	color.Cyan("|             MENU              |")
	fmt.Println("+-------------------------------+")
	color.Blue("|👕 Informations personnage [P] |")
	color.Blue("|🎒 Accéder à l’inventaire [I]  |")
	color.Green("|🌟 Potion de soin [S]          |")
	color.HiGreen("|☠️  Potion de poison [U]        |")
	color.HiBlack("|💶 Magasin [M]                 |")
	color.HiBlack("|⚔️  Forgeron [F]                |")
	fmt.Println("|                               |")
	color.Red("|❌ Quitter le jeu [Exit]       |")
	fmt.Println("+-------------------------------+")

	// Remplacement du second "MENU" par "Votre choix ?"
	color.Yellow("\nVotre choix ? ")
	var choice string
	fmt.Scanln(&choice)


				switch choice {
		case "P":
			displayInfo(*c1)
		case "I":
			c1.accessInventory()
		case "S":
			c1.TakePot()
		case "U":
			c1.PoisonPot()
		case "M":
    		Merchant(c1)
		case "F":
			color.HiBlack("Forgeron : Pas encore codé")
		case "Exit":
			color.Red("Fermeture du jeu...")
			return
		default:
			color.Red("Choix non reconnu")
        }
		c1.IsDead()
	}
}

// Tâche 8 : Vérification de la mort du personnage
func (personnage *Character) IsDead() {
	if personnage.PV <= 0 {
		color.HiRed("%s a succombé à ses blessures ! ⚰️", personnage.Name)
		// Résurrection avec 50% des PV
		personnage.PV = 100
		color.Green("%s est ressuscité avec %d/%d PV ! ✨",
			personnage.Name, personnage.PV, personnage.Max_PV)
	}
}

// Fonction main
func main() {
	startGame()

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

	// Étape 2 : lancements 
	menu(&c1)
	c1.IsDead()
}