// Ce fichier contient toutes les fonctions nécessaires au service de vente (magasin + forge) et à leur fonctionnement  
// Le paquet de la librairie où sont stockées les fonctions 

package librairies

import (
	"fmt"
	"github.com/fatih/color"
)

// ------------- Structures ----------------


// Cette structure définit les tarifs et effets des objets du magasin
type ShopItem struct {
	Name       string
	PriceRubis int
	PriceDiam  int
	EffectPV   int // indicatif (affiché seulement), l’effet n’est pas appliqué dans le shop
}

// Cette structure concerne les objets de la forge et leurs matériaux
type ForgeItem struct {
	Name      string
	Materials map[string]int
	Gender    string // "m", "f", "fp"
}

// ------------- Données du magasin ----------------

// Stock global du magasin (persiste durant toute la partie)
var shopItems = []ShopItem{
	{"Arrow", 10, 0, 0},
	{"Arrow", 10, 0, 0},
	{"Arrow", 10, 0, 0},
	{"Arrow", 10, 0, 0},
	{"Bow", 5, 0, 0},
	{"Cuir", 5, 0, 0},
	{"Divine Venison", 25, 0, 25},
	{"Divine Venison", 25, 0, 25},
	{"Fairy", 50, 0, 55},
	{"Fairy", 50, 0, 55},
	{"Lingot", 3, 0, 0},
	{"Lingot", 3, 0, 0},
	{"Master Sword", 0, 2, 0},
	{"Miasme", 25, 0, -45},
	{"Miasme", 25, 0, -45},
	{"Upgrade kit", 10, 0, 0},
	{"Upgrade kit", 10, 0, 0},
	{"Upgrade kit", 10, 0, 0},
}

// ------------- Fonctions du magasin ----------------

// Afficher les objets du magasin
func printShop(items []ShopItem) {
	fmt.Println("+----+-----------------+-------------+--------+")
	fmt.Printf("| %-2s | %-15s | %-11s | %-6s |\n", "N°", "Items", "Prix", "PV")
	fmt.Println("+----+-----------------+-------------+--------+")

	for i, item := range items {
		// Prix affiché
		price := ""
		if item.PriceRubis > 0 {
			price = fmt.Sprintf("%d rubis", item.PriceRubis)
		} else if item.PriceDiam > 0 {
			price = fmt.Sprintf("%d diamants", item.PriceDiam)
		} else {
			price = "gratuit"
		}

		// PV affichés (info uniquement)
		pv := "-"
		if item.EffectPV != 0 {
			pv = fmt.Sprintf("%+d", item.EffectPV)
		}

		fmt.Printf("| %-2d | %-15s | %-11s | %-6s |\n", i+1, item.Name, price, pv)
	}

	fmt.Println("+----+-----------------+-------------+--------+")
	fmt.Printf("| %-2s | %-15s | %-11s | %-6s |\n", "0", "Quitter", "", "")
	fmt.Println("+----+-----------------+-------------+--------+")
}

// Interaction avec le marchand
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

	// Cadeau diamant la première fois
	if !personnage.HasReceivedDiamond {
		personnage.AddInventory("Diamant")
		personnage.HasReceivedDiamond = true
		fmt.Println("\n Vous avez reçu un 💎 gratuit ! \n")
	}

	for {
		if len(shopItems) == 0 {
			fmt.Println("Le magasin est vide.")
			return
		}

		printShop(shopItems)
		fmt.Println("\nRubis disponibles :", personnage.Rubis)

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

		// Vérification de l’argent
		if selectedItem.PriceRubis > personnage.Rubis {
			fmt.Println("Pas assez de Rubis !")
			continue
		}
		if selectedItem.PriceDiam > personnage.CountItem("Diamant") {
			fmt.Println("Pas assez de Diamants !")
			continue
		}

		// Paiement
		personnage.Rubis -= selectedItem.PriceRubis
		for i := 0; i < selectedItem.PriceDiam; i++ {
			personnage.RemoveItem("Diamant")
		}

		// Cas spécial upgrade kit
		if selectedItem.Name == "Upgrade kit" {
			personnage.UpgradeInventorySlot()
		} else {
			// Tous les autres vont dans l’inventaire
			personnage.AddInventory(selectedItem.Name)
		}

		// Retirer l’item du shop
		shopItems = append(shopItems[:choix-1], shopItems[choix:]...)
	}
}

// ------------- FORGE ----------------

// Affiche le menu du forgeron
func printForgeMenu(items []ForgeItem) {
	color.Cyan("+----------------------------------------+")
	color.Cyan("|               Forgeron                 |")
	color.Cyan("+----------------------------------------+")

	for i, it := range items {
		// Construire la liste des matériaux
		matList := ""
		first := true
		for mat, qty := range it.Materials {
			if !first {
				matList += ", "
			}
			matList += fmt.Sprintf("%dx %s", qty, mat)
			first = false
		}

		// Affichage aligné
		fmt.Printf("| %d) %-15s | %s\n", i+1, it.Name, matList)
	}

	color.Cyan("| 0) Quitter                             |")
	color.Cyan("+----------------------------------------+")
}

// Génère le message de forge correct selon le genre/nombre
func forgeMessage(name, gender string) string {
	switch gender {
	case "f":
		return fmt.Sprintf("%s a été forgée et ajoutée à votre équipement !", name)
	case "fp":
		return fmt.Sprintf("%s ont été forgées et ajoutées à votre équipement !", name)
	default: // masculin singulier
		return fmt.Sprintf("%s a été forgé et ajouté à votre équipement !", name)
	}
}

// Fonction principale de la forge
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

	// Les items du forgeron (forgeables une seule fois chacun)
	forgeItems := []ForgeItem{
		{"Casque de garde", map[string]int{"Lingot": 1}, "m"},
		{"Tunique royale", map[string]int{"Lingot": 1, "Tissu royal": 1}, "f"},
		{"Bottes", map[string]int{"Cuir": 1}, "fp"},
	}

	for {
		printForgeMenu(forgeItems)
		fmt.Println("\nInventaire actuel :", personnage.Inventory)

		var choix int
		fmt.Print("Votre choix : ")
		fmt.Scan(&choix)

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

		// Vérifier si le joueur a tous les matériaux nécessaires
		for mat, qty := range selectedItem.Materials {
			if personnage.CountItem(mat) < qty {
				fmt.Printf("Il vous manque %dx %s pour forger %s.\n", qty, mat, selectedItem.Name)
				canForge = false
			}
		}
		if !canForge {
			continue
		}

		// Retirer les matériaux de l'inventaire
		for mat, qty := range selectedItem.Materials {
			for i := 0; i < qty; i++ {
				personnage.RemoveItem(mat)
			}
		}

		// Ajouter l'item forgé à l'équipement
		personnage.AddEquipment(selectedItem.Name)
		fmt.Println(forgeMessage(selectedItem.Name, selectedItem.Gender))

		// Supprimer l'item de la forge (forgeable une seule fois)
		forgeItems = append(forgeItems[:choix-1], forgeItems[choix:]...)

		if len(forgeItems) == 0 {
			fmt.Println("Le forgeron n'a plus d'items. Retour au menu principal.")
			return
		}
	}
}
