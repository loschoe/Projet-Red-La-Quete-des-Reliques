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

// -------- STRUCTURES --------
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

type ShopItem struct {
	Name       string
	PriceRubis int
	PriceDiam  int
	EffectPV   int
}

type ForgeItem struct {
	Name      string
	Materials map[string]int // ex: {"Lingot": 1, "Cuir": 4}
	EffectPV  int
}

// -------- INITIALISATION --------
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
		Rubis:     10, // commence avec 10 rubis
	}
}

// -------- INVENTAIRE --------
func displayInfo(c Character) {
	fmt.Printf("\nName : %s\nClasse : %s\nLevel : %d\nPV : %d/%d\nInventory : %v\n",
		c.Name, c.Classe, c.Level, c.PV, c.Max_PV, c.Inventory)
}

func (personnage *Character) accessInventory() {
	fmt.Println("\nInventaire du personnage :")
	vide := true
	for i, item := range personnage.Inventory {
		if item == "" || item == "..." {
			fmt.Printf("%d. [vide]\n", i+1)
		} else {
			fmt.Printf("%d. %s\n", i+1, item)
			vide = false
		}
	}
	if vide {
		fmt.Println("Inventaire vide.")
	}
	fmt.Println("Rubis disponibles :", personnage.Rubis)
	fmt.Println()
}

func (personnage *Character) AddInventory(item string) {
	if personnage.IsInventoryFull() {
		fmt.Println("Inventaire plein ! Impossible d'ajouter", item)
		return
	}
	for i := 0; i < len(personnage.Inventory); i++ {
		if personnage.Inventory[i] == "" || personnage.Inventory[i] == "..." {
			personnage.Inventory[i] = item
			fmt.Println(item, "a été ajouté à l'inventaire.")
			return
		}
	}
}

func (personnage *Character) RemoveItem(item string) {
	for idx, i := range personnage.Inventory {
		if i == item {
			personnage.Inventory[idx] = ""
			fmt.Println(item, "a été retiré de l'inventaire.")
			return
		}
	}
}

func (personnage *Character) CountItem(item string) int {
	count := 0
	for _, i := range personnage.Inventory {
		if i == item {
			count++
		}
	}
	return count
}

func (personnage *Character) RemoveItemAt(index int) {
	for j := index; j < len(personnage.Inventory)-1; j++ {
		personnage.Inventory[j] = personnage.Inventory[j+1]
	}
	personnage.Inventory[len(personnage.Inventory)-1] = ""
}

func (personnage *Character) IsInventoryFull() bool {
	count := 0
	for _, item := range personnage.Inventory {
		if item != "" && item != "..." {
			count++
		}
	}
	return count >= 10
}

// -------- SHOP --------
func printShop(shopItems []ShopItem) {
	fmt.Println("+-----------------+-------------+")
	fmt.Printf("| %-15s | %-11s |\n", "Items", "Prix")
	fmt.Println("+-----------------+-------------+")

	for _, item := range shopItems {
		price := ""
		if item.PriceRubis > 0 {
			price = fmt.Sprintf("%d rubis", item.PriceRubis)
		} else if item.PriceDiam > 0 {
			price = fmt.Sprintf("%d diamants", item.PriceDiam)
		} else {
			price = "gratuit"
		}
		fmt.Printf("| %-15s | %-11s |\n", item.Name, price)
	}

	fmt.Println("+-----------------+-------------+")
	fmt.Printf("| %-15s | %-11s |\n", "Quitter", "Press 0")
	fmt.Println("+-----------------+-------------+")
}

func Merchant(personnage *Character) {
	shopArt := `
                              _       
  /\/\   __ _  __ _  __ _ ___(_)_ __  
 /    \ / _' |/ _' |/ _' / __| | '_ \ 
/ /\/\ \ (_| | (_| | (_| \__ \ | | | |
\/    \/\__,_|\__, |\__,_|___/_|_| |_|
              |___/                   
`
	color.Red("%s\n", shopArt)

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
		fmt.Println("Vous avez reçu un Diamant gratuit !")
	}

	for {
		printShop(shopItems)

		fmt.Println("\nRubis disponibles :", personnage.Rubis)

		var choix int
		fmt.Print("Votre choix : ")
		fmt.Scanln(&choix)

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
			fmt.Println("Pas assez de Rubis !")
			continue
		}
		if selectedItem.PriceDiam > personnage.CountItem("Diamant") {
			fmt.Println("Pas assez de Diamants !")
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

		shopItems = append(shopItems[:choix-1], shopItems[choix:]...)
	}
}

// -------- FORGE --------
func printForge(items []ForgeItem) {
	fmt.Println("\n--- Forgeron ---")
	for i, it := range items {
		fmt.Printf("%d) %s (nécessite : ", i+1, it.Name)
		first := true
		for mat, qty := range it.Materials {
			if !first {
				fmt.Print(", ")
			}
			fmt.Printf("%dx %s", qty, mat)
			first = false
		}
		fmt.Println(")")
	}
	fmt.Println("0) Quitter")
}

func Forge(personnage *Character) {
	forgeArt := `
  ______                     
 |  ____|                    
 | |__ ___  _ __ __ _  ___  
 |  __/ _ \| '__/ _' |/ _ \ 
 | | | (_) | | | (_| |  __/ 
 |_|  \___/|_|  \__, |\___| 
                 __/ |      
                |___/       
`
	color.HiBlack("%s\n", forgeArt)

	forgeItems := []ForgeItem{
		{"Casque de garde", map[string]int{"Lingot": 1}, 0},
		{"Tunique royale", map[string]int{"Lingot": 1, "Tissu royal": 1}, 0},
		{"Bottes", map[string]int{"4x Cuir": 1}, 0},
	}

	for {
		printForge(forgeItems)
		fmt.Println("\nInventaire :", personnage.Inventory)

		var choix int
		fmt.Print("Votre choix : ")
		fmt.Scanln(&choix)

		if choix == 0 {
			fmt.Println("Au revoir !")
			return
		}

		if choix < 1 || choix > len(forgeItems) {
			fmt.Println("Choix invalide.")
			continue
		}

		selectedItem := forgeItems[choix-1]
		canForge := true
		for mat, qty := range selectedItem.Materials {
			if personnage.CountItem(mat) < qty {
				fmt.Printf("Il vous manque %dx %s pour forger %s.\n", qty, mat, selectedItem.Name)
				canForge = false
			}
		}
		if !canForge {
			continue
		}

		for mat, qty := range selectedItem.Materials {
			for i := 0; i < qty; i++ {
				personnage.RemoveItem(mat)
			}
		}

		personnage.AddInventory(selectedItem.Name)
		fmt.Println(selectedItem.Name, "a été forgé et ajouté à votre inventaire !")

		forgeItems = append(forgeItems[:choix-1], forgeItems[choix:]...)

		if len(forgeItems) == 0 {
			fmt.Println("Le forgeron n'a plus d'items. Retour au menu principal.")
			return
		}
	}
}

// -------- POTIONS --------
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
	fmt.Println("Aucune Fée disponible.")
}

func (personnage *Character) PoisonPot() {
	for i, item := range personnage.Inventory {
		if item == "Miasme" {
			fmt.Println(personnage.Name, "utilise un miasme !")
			for j := 1; j <= 3; j++ {
				time.Sleep(1 * time.Second)
				personnage.PV -= 15
				if personnage.PV < 0 {
					personnage.PV = 0
				}
				fmt.Printf("Après %d seconde(s) : %d / %d PV\n", j, personnage.PV, personnage.Max_PV)
			}
			fmt.Println("Le miasme n’a plus d’effet")
			personnage.RemoveItemAt(i)
			return
		}
	}
	fmt.Println("Aucun Miasme disponible.")
}

// -------- MENU --------
func menu(c1 *Character) {
	for {
		fmt.Println("+-------------------------------+")
		color.Cyan("|             MENU              |")
		fmt.Println("+-------------------------------+")
		color.Blue("|👕 Infos personnage [P]        |")
		color.Blue("|🎒 Inventaire [I]              |")
		color.Green("|🌟 Potion de soin [S]          |")
		color.HiGreen("|☠️  Potion de poison [U]        |")
		color.HiBlack("|💶 Magasin [M]                 |")
		color.HiBlack("|⚔️  Forgeron [F]                |")
		fmt.Println("|                               |")
		color.Red("|❌ Quitter le jeu [Exit]       |")
		fmt.Println("+-------------------------------+")

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
			Forge(c1)
		case "Exit":
			color.Red("Fermeture du jeu...")
			return
		default:
			color.Red("Choix non reconnu")
		}
		c1.IsDead()
	}
}

// -------- SYSTEME DE MORT --------
func (personnage *Character) IsDead() {
	if personnage.PV <= 0 {
		color.HiRed("%s est mort ! ⚰️", personnage.Name)
		personnage.PV = personnage.Max_PV / 2
		color.Green("%s est ressuscité avec %d/%d PV ! ✨", personnage.Name, personnage.PV, personnage.Max_PV)
	}
}

// -------- MAIN --------
func main() {
	startGame()

	var inventory [10]string
	inventory[0] = "Fairy"
	inventory[1] = "Lingot"
	inventory[2] = "4x Cuir"
	inventory[3] = "Tissu royal"
	inventory[4] = "Lingot"
	inventory[5] = "Tissu royal"

	c1 := initCharacter("Link", "Hylien", 1, 500, 100, inventory)

	for i, item := range c1.Inventory {
		if item == "" {
			c1.Inventory[i] = "..."
		}
	}

	menu(&c1)
}
