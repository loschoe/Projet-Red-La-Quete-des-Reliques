package main

import "fmt"



// Supprimer le premier exemplaire d'un item précis
func (personnage *Character) RemoveItem(item string) {
	for idx, i := range personnage.Inventory {
		if i == item {
			personnage.Inventory[idx] = ""
			fmt.Println(item, "a été retiré de l'inventaire.")
			return
		}
	}
	fmt.Println("Aucun", item, "n'a été trouvé dans l'inventaire.")
}

// Afficher l'inventaire
func (personnage *Character) accessInventory() {
	fmt.Println("\nInventaire du personnage :")
	vide := true
	for i, item := range personnage.Inventory {
		if item != "" {
			fmt.Printf("%d. %s (utilisable)\n", i+1, item)
			vide = false
		} else {
			fmt.Printf("%d. [vide]\n", i+1)
		}
	}
	if vide {
		fmt.Println("L'inventaire de votre personnage est vide.")
	}
	fmt.Println("Rubis disponibles :", personnage.Rubis)
	fmt.Println()
}

// Marchand complet
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
		fmt.Println("Vous avez reçu un Diamant gratuit !")
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
			fmt.Println("Vous n'avez pas assez de Rubis pour cet item !")
			continue
		}
		if selectedItem.PriceDiam > personnage.CountItem("Diamant") {
			fmt.Println("Vous n'avez pas assez de Diamants pour cet item !")
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

// Menu principal
func main() {
	var emptyInventory [10]string
	player := Character{
		Name:               "Link",
		Classe:             "Héros",
		Level:              1,
		Max_PV:             100,
		PV:                 100,
		Inventory:          emptyInventory,
		HasReceivedDiamond: false,
		Rubis:              100,
	}

	for {
		fmt.Println("\n===== MENU =====")
		fmt.Println("1. Voir l'inventaire")
		fmt.Println("2. Aller chez le marchand")
		fmt.Println("0. Quitter")
		fmt.Print("Votre choix : ")

		var choix int
		fmt.Scan(&choix)

		switch choix {
		case 0:
			fmt.Println("Au revoir !")
			return
		case 1:
			player.accessInventory()
		case 2:
			Merchant(&player)
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
