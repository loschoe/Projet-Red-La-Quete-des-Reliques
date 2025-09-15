package librairies

import (
	"fmt"
	"github.com/fatih/color"
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
		{"Master Sword", 0, 2, 0},
		{"PoisonPot", 25, 0, 0},
		{"Divine Venison", 50, 0, 25},
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
