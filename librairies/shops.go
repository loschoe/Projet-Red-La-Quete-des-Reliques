package librairies

import (
	"fmt"
	"github.com/fatih/color"
	"sort"
)

type ShopItem struct {
	Name       string
	PriceRubis int
	PriceDiam  int
	EffectPV   int
}

type ForgeItem struct {
	Name      string
	Materials map[string]int
	EffectPV  int
}

func printShop(shopItems []ShopItem) {
	// Copier et trier les items par nom
	sortedItems := make([]ShopItem, len(shopItems))
	copy(sortedItems, shopItems)
	sort.Slice(sortedItems, func(i, j int) bool {
		return sortedItems[i].Name < sortedItems[j].Name
	})

	fmt.Println("+----+-----------------+-------------+--------+")
	fmt.Printf("| %-2s | %-15s | %-11s | %-6s |\n", "N°", "Items", "Prix", "PV")
	fmt.Println("+----+-----------------+-------------+--------+")

	for i, item := range sortedItems {
		// Détermination du prix
		price := ""
		if item.PriceRubis > 0 {
			price = fmt.Sprintf("%d rubis", item.PriceRubis)
		} else if item.PriceDiam > 0 {
			price = fmt.Sprintf("%d diamants", item.PriceDiam)
		} else {
			price = "gratuit"
		}

		// Affichage des PV (positifs ou négatifs)
		pv := "-"
		if item.EffectPV != 0 {
			pv = fmt.Sprintf("%+d", item.EffectPV)
		}

		fmt.Printf("| %-2d | %-15s | %-11s | %-6s |\n", i+1, item.Name, price, pv)

		// Bordure après chaque groupe d'items identiques
		if i == len(sortedItems)-1 || sortedItems[i].Name != sortedItems[i+1].Name {
			fmt.Println("+----+-----------------+-------------+--------+")
		}
	}

	// Ligne pour quitter
	fmt.Printf("| %-2s | %-15s | %-11s | %-6s |\n", "0", "Quitter", "", "")
	fmt.Println("+----+-----------------+-------------+--------+")
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
		{"Arrow", 10, 0, 0},
		{"Arrow", 10, 0, 0},
		{"Arrow", 10, 0, 0},
		{"Arrow", 10, 0, 0},
		{"Master Sword", 0, 2, 0},
		{"Bow", 5, 0, 0},
		{"Miasme", 25, 0, -45},
		{"Miasme", 25, 0, -45},
		{"Fairy", 50, 0, 55},
		{"Fairy", 50, 0, 55},
		{"Divine Venison", 25, 0, 25},
		{"Divine Venison", 25, 0, 25},
		{"Lingot", 3, 0, 0},
		{"Lingot", 3, 0, 0},
		{"Cuir", 5, 0, 0},
		{"Upgrade kit", 10,0,0},
		{"Upgrade kit", 10,0,0},
		{"Upgrade kit", 10,0,0},
	}

	if !personnage.HasReceivedDiamond {
		personnage.AddInventory("Diamant")
		personnage.HasReceivedDiamond = true
		fmt.Println("\n Vous avez reçu un 💎 gratuit ! \n")
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

		if selectedItem.Name == "Upgrade kit" {
			personnage.Rubis -= selectedItem.PriceRubis
			personnage.UpgradeInventorySlot() // on appelle la fonction
			// On supprime l’item du shop pour éviter qu’il soit affiché à l’infini
			shopItems = append(shopItems[:choix-1], shopItems[choix:]...)
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
		{"Bottes", map[string]int{"Cuir": 1}, 0},
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
