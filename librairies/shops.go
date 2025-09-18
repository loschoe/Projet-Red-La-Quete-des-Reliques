// Ce fichier contient toutes les fonctions nécessaires au service de vente (magasin + forge) et à leur fonctionnement  
// Le paquet de la librairie où sont stockées les fonctions 

package librairies

import (
	"fmt"
	"github.com/fatih/color"
	"strings"
)

// ------------- Structures ----------------


// Cette structure définit les tarifs et effets des objets du magasin
type ShopItem struct {
	Name       string
	PriceRubis int
	PriceDiam  int
	EffectPV   int // indicatif (affiché seulement), l’effet n’est pas appliqué dans le shop
	Quantity int 
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
	{"Arrow", 15, 0, 0, 5},
	{"Bow", 20, 0, 0, 1},
	{"Cuir", 10, 0, 0, 1},
	{"Divine Venison", 30, 0, 25, 2},
	{"Fairy", 60, 0, 55, 10},
	{"Lingot", 20, 0, 0, 2},
	{"Master Sword", 0, 2, 0, 1},
	{"Miasme", 25, 0, -45, 5},
	{"Upgrade kit", 12, 0, 0, 3},
	{"Zelda Book", 100, 0, 10, 1},
}

// ------------- Fonctions du magasin ----------------

// Afficher les objets du magasin
func printShop(items []ShopItem) {
	fmt.Println("+----+-----------------+-------------+--------+-------+")
	fmt.Printf("| %-2s | %-15s | %-11s | %-6s | %-5s |\n", "N°", "Items", "Prix", "PV", "Qté")
	fmt.Println("+----+-----------------+-------------+--------+-------+")

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

		// Ligne d’un item
		fmt.Printf("| %-2d | %-15s | %-11s | %-6s | %-5d |\n",
			i+1, item.Name, price, pv, item.Quantity)
	}

	fmt.Println("+----+-----------------+-------------+--------+-------+")
	fmt.Printf("| %-2s | %-15s | %-11s | %-6s | %-5s |\n", "0", "Quitter", "", "", "")
	fmt.Println("+----+-----------------+-------------+--------+-------+")
}


// Interaction avec le marchand
func Merchant(personnage *Character) {
	shopArt := `
                              _       
  /\/\   __ _  __ _  __ _ ___(_)_ __  
 /    \ / _' |/ _' |/ _' |/ _' / __| | '_ \ 
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
			color.Red("Choix invalide.")
			continue
		}

		selectedItem := &shopItems[choix-1] // pointeur pour modifier directement Quantity

		// Vérification de l’argent
		if selectedItem.PriceRubis > personnage.Rubis {
			color.Red("Pas assez de Rubis !")
			continue
		}
		if selectedItem.PriceDiam > personnage.CountItem("Diamant") {
			color.Red("Pas assez de Diamants !")
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
			personnage.AddInventory(selectedItem.Name)
		}

		// Décrémenter la quantité
		selectedItem.Quantity--

		// Retirer l’item si stock épuisé
		if selectedItem.Quantity <= 0 {
			shopItems = append(shopItems[:choix-1], shopItems[choix:]...)
		}
	}
}

// ---------------- FORGE ----------------
// Items globaux du forgeron (forgeables une seule fois)
var forgeItems = []ForgeItem{
	{"Casque de garde", map[string]int{"Lingot": 1}, "m"},
	{"Tunique royale", map[string]int{"Lingot": 1, "Tissu royal": 1}, "f"},
	{"Bottes", map[string]int{"Cuir": 1}, "fp"},
}

// Affiche le menu du forgeron
func printForgeMenu(items []ForgeItem) {
	totalWidth := 48 // largeur de la zone interne (entre les |)

	color.Cyan("+--------------------------------------------------+")
	color.Cyan("|                    Forgeron                      |")
	color.Cyan("+--------------------------------------------------+")

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

		// Construire la ligne (numéro + nom + matériaux)
		line := fmt.Sprintf("%d) %-15s | %s", i+1, it.Name, matList)

		// Compléter avec des espaces pour atteindre la largeur interne
		if len(line) < totalWidth {
			line += strings.Repeat(" ", totalWidth-len(line))
		}

		color.Yellow("| " + line + " |")
	}

	// Ligne Quitter (même traitement que les autres lignes)
	quitLine := "0) Quitter"
	if len(quitLine) < totalWidth {
		quitLine += strings.Repeat(" ", totalWidth-len(quitLine))
	}
	color.Cyan("| " + quitLine + " |")

	color.Cyan("+--------------------------------------------------+")
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

	for {
		if len(forgeItems) == 0 {
			fmt.Println("Le forgeron n'a plus d'items. Retour au menu principal.")
			return
		}

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
	}
}

