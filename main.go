package main

import (
	"fmt"
	"github.com/fatih/color"
	"time"
)

// ------------------ DÉMARRAGE ------------------
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

	  Appuyez sur Entrée pour commencer !!
	`
	color.Cyan("%s\n", asciiArt)
	color.Red("%s\n", introText)
	fmt.Scanln()
}

// ------------------ STRUCTURES ------------------
type Character struct {
	Name               string
	Classe             string
	Level              int
	Max_PV             int
	PV                 int
	Inventory          []string
	InventoryCapacity  int
	InventoryUpgrades  int
	HasReceivedDiamond bool
	Rubis              int
}

type ShopItem struct {
	Name       string
	PriceRubis int
	PriceDiam  int
	EffectPV   int
}

// ------------------ INITIALISATION ------------------
func initCharacter(name string, classe string, level, max_pv, pv int, inventory []string) Character {
	if pv > max_pv {
		pv = max_pv
	}
	return Character{
		Name:              name,
		Classe:            classe,
		Level:             level,
		Max_PV:            max_pv,
		PV:                pv,
		Inventory:         inventory,
		InventoryCapacity: 10,
		Rubis:             1000,
	}
}

// ------------------ AFFICHAGE ------------------
func displayInfo(c Character) {
	fmt.Printf("\nNom : %s\nClasse : %s\nLevel : %d\nPV : %d/%d\nInventaire : %v\n",
		c.Name, c.Classe, c.Level, c.PV, c.Max_PV, c.Inventory)
}

func (p *Character) accessInventory() {
	fmt.Println("\nInventaire du personnage :")
	vide := true
	for i := 0; i < p.InventoryCapacity; i++ {
		if i >= len(p.Inventory) || p.Inventory[i] == "" {
			fmt.Printf("%d. [vide]\n", i+1)
		} else {
			fmt.Printf("%d. %s (utilisable)\n", i+1, p.Inventory[i])
			vide = false
		}
	}
	if vide {
		fmt.Println("L'inventaire de votre personnage est vide.")
	}
	fmt.Println("Rubis disponibles :", p.Rubis)
	fmt.Println()
}

// ------------------ INVENTAIRE ------------------
func (p *Character) AddInventory(item string) {
	if len(p.Inventory) >= p.InventoryCapacity {
		fmt.Println("Inventaire plein ! Impossible d'ajouter", item)
		return
	}
	p.Inventory = append(p.Inventory, item)
	fmt.Println(item, "a été ajouté à l'inventaire avec succès !")
}

func (p *Character) RemoveItem(item string) {
	for idx, i := range p.Inventory {
		if i == item {
			p.Inventory[idx] = ""
			fmt.Println(item, "a été retiré de l'inventaire.")
			return
		}
	}
	fmt.Println("Aucun(e)", item, "n'a été trouvé dans l'inventaire.")
}

func (p *Character) RemoveItemAt(index int) {
	for j := index; j < len(p.Inventory)-1; j++ {
		p.Inventory[j] = p.Inventory[j+1]
	}
	p.Inventory[len(p.Inventory)-1] = ""
}

func (p *Character) CountItem(item string) int {
	count := 0
	for _, i := range p.Inventory {
		if i == item {
			count++
		}
	}
	return count
}

// ------------------ POTIONS ------------------
func (p *Character) TakePot() {
	for i, item := range p.Inventory {
		if item == "Fairy" {
			p.PV += 50
			if p.PV > p.Max_PV {
				p.PV = p.Max_PV
			}
			fmt.Println(p.Name, "utilise une Fée ! PV =", p.PV, "/", p.Max_PV)
			p.RemoveItemAt(i)
			return
		}
	}
	fmt.Println("Aucune Fée n'est disponible dans l'inventaire.")
}

func (p *Character) PoisonPot() {
	for i, item := range p.Inventory {
		if item == "Miasme" {
			fmt.Println(p.Name, "utilise un Miasme !")
			for j := 1; j <= 3; j++ {
				time.Sleep(1 * time.Second)
				p.PV -= 15
				if p.PV < 0 {
					p.PV = 0
				}
				fmt.Printf("Après %d seconde(s) : %d / %d PV\n", j, p.PV, p.Max_PV)
			}
			fmt.Println("Le Miasme n’a plus d’effet")
			p.RemoveItemAt(i)
			return
		}
	}
	fmt.Println("Aucun Miasme n'est disponible dans l'inventaire.")
}

// ------------------ UPGRADE INVENTAIRE ------------------
func (p *Character) UpgradeInventorySlot() {
	if p.InventoryUpgrades >= 3 {
		fmt.Println("Vous ne pouvez plus augmenter l’inventaire.")
		return
	}
	p.InventoryCapacity += 5 // augmentation de 5 slots
	p.InventoryUpgrades++
	fmt.Printf("L’inventaire a été augmenté ! Nouvelle capacité : %d (utilisations : %d/3)\n",
		p.InventoryCapacity, p.InventoryUpgrades)
}

// ------------------ SHOP ------------------
func printShop(shopItems []ShopItem) {
	fmt.Println("+---------------------------+----------------+")
	fmt.Printf("| %-25s | %-15s |\n", "Items", "Prix")
	fmt.Println("+---------------------------+----------------+")
	for _, item := range shopItems {
		price := ""
		if item.PriceRubis > 0 {
			price = fmt.Sprintf("%d rubis", item.PriceRubis)
		} else if item.PriceDiam > 0 {
			price = fmt.Sprintf("%d diamants", item.PriceDiam)
		} else {
			price = "gratuit"
		}
		fmt.Printf("| %-25s | %-15s |\n", item.Name, price)
	}
	fmt.Println("+---------------------------+----------------+")
	fmt.Printf("| %-25s | %-15s |\n", "Quitter", "Press 0")
	fmt.Println("+---------------------------+----------------+")
}

func Merchant(p *Character) {
	color.Red("\nBienvenue au Marchand !\n")

	shopItems := []ShopItem{
		{"5X Arrow", 10, 0, 0},
		{"5X Arrow", 10, 0, 0},
		{"Master Sword", 0, 3, 0},
		{"PoisonPot", 25, 0, 0},
		{"Divine Venison", 50, 0, 25},
		{"Augmentation d'inventaire", 30, 0, 0},
		{"Augmentation d'inventaire", 30, 0, 0},
		{"Augmentation d'inventaire", 30, 0, 0},
		{"Iron ingot", 5, 0, 0},
		{"4X Leather", 15, 0, 0},
	}

	if !p.HasReceivedDiamond {
		p.AddInventory("Diamant")
		p.HasReceivedDiamond = true
		fmt.Println("Vous avez reçu un Diamant gratuit !")
	}

	for {
		printShop(shopItems)
		fmt.Println("Rubis disponibles :", p.Rubis)

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

		if selectedItem.PriceRubis > p.Rubis {
			fmt.Println("Pas assez de Rubis pour acheter cet item !")
			continue
		}
		if selectedItem.PriceDiam > p.CountItem("Diamant") {
			fmt.Println("Pas assez de Diamants pour obtenir cet item !")
			continue
		}

		p.Rubis -= selectedItem.PriceRubis
		for i := 0; i < selectedItem.PriceDiam; i++ {
			p.RemoveItem("Diamant")
		}

		if selectedItem.Name == "Augmentation d'inventaire" {
			p.UpgradeInventorySlot()
		} else {
			p.AddInventory(selectedItem.Name)
		}

		if selectedItem.EffectPV > 0 {
			p.PV += selectedItem.EffectPV
			if p.PV > p.Max_PV {
				p.PV = p.Max_PV
			}
			fmt.Printf("%s vous rend %d PV ! PV actuel : %d/%d\n", selectedItem.Name, selectedItem.EffectPV, p.PV, p.Max_PV)
		}

		fmt.Println(selectedItem.Name, "a été ajouté à votre inventaire !")
		shopItems = append(shopItems[:choix-1], shopItems[choix:]...)
		if len(shopItems) == 0 {
			fmt.Println("Le marchand n'a plus d'items à vendre. Retour au menu principal.")
			return
		}
	}
}

// ------------------ MENU ------------------
func menu(c1 *Character) {
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
			color.Red("Choix non reconnu.")
		}
		c1.IsDead()
	}
}

// ------------------ MORT ------------------
func (p *Character) IsDead() {
	if p.PV <= 0 {
		color.HiRed("%s a succombé à ses blessures ! ⚰️\n", p.Name)
		p.PV = p.Max_PV / 2
		color.Green("%s est ressuscité avec %d/%d PV ! ✨\n", p.Name, p.PV, p.Max_PV)
	}
}

// ------------------ MAIN ------------------
func main() {
	startGame()
	inventory := []string{"Fairy", "Fairy", "Fairy"}
	c1 := initCharacter("Link", "Hylien", 1, 500, 100, inventory)
	menu(&c1)
	c1.IsDead()
}
