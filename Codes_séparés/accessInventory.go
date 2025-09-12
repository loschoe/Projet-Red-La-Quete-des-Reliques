package main

import "fmt"

// Pour génériser votre code, vous pouvez créer des fonctions telles que addInventory et removeInventory pour
// gérer l’ajout et le retrait d’item de l’inventaire

// Ajoute un item dans la première case libre
func (personnage *Character) AddInventory(item string) {
	for i := 0; i < len(personnage.Inventory); i++ {
		if personnage.Inventory[i] == "" {
			personnage.Inventory[i] = item
			fmt.Println(item, "a été ajouté à l'inventaire avec succès !")
			return
		}
	}
	fmt.Println("Inventaire plein ! Impossible d'ajouter", item)
}

// Retire un item d'une position donnée (1-10)
func (personnage *Character) RemoveInventory(pos int) {
	if pos < 1 || pos > 10 {
		fmt.Println("Position invalide.")
		return
	}
	if personnage.Inventory[pos-1] == "" {
		fmt.Println("Cette case est déjà vide.")
		return
	}
	fmt.Println(personnage.Inventory[pos-1], "a été retiré de l'inventaire.")
	personnage.Inventory[pos-1] = ""
}

// Compte combien d'exemplaires d'un item précis sont dans l'inventaire
func (personnage *Character) CountItem(item string) int {
	count := 0
	for _, i := range personnage.Inventory {
		if i == item {
			count++
		}
	}
	return count
}


// Créez la fonction accessInventory qui permet d’afficher tous les items présents dans l’inventaire du personnage qui seront utilisables par la suite
func (personnage *Character) accessInventory() {
	fmt.Println("Inventaire du personnage :")

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
}

func Merchant(personnage *Character) {
	fmt.Println("Bienvenue dans ma boutique ! Voici ce que j'ai en stock :")
	shopItems := [5]string{"5X Arrow", "5X Arrow", "Master Sword", "PoisonPot", "Divine Venison"}

	// Affichage des items
	for i, item := range shopItems {
		fmt.Printf("%d. %s\n", i+1, item)
	}
	fmt.Println("Pour débloquer le Bouclier d'Hylia vous devez terminer 5 entraînements.")
	fmt.Println("En attendant, je vous offre gratuitement votre premier diamant ! Il vous servira pour acheter l'épée de légende.")

	// Ajouter le diams, gérer le système de prix/d'achats 
	// Le diamant est donné QUE lors de la première visite
	
	personnage.AddInventory("Diamant")
	fmt.Println("Vous avez reçu un Diamant !")

	



	fmt.Println("0. Quitter")

	// Choix du joueur
	var choix int
	fmt.Print("Votre choix : ")
	fmt.Scan(&choix)

	if choix == 0 {
		fmt.Println("Au revoir !")
		return
	}

	if choix < 0 || choix > len(shopItems) {
		fmt.Println("Choix invalide.")
		return
	}

	// Ajouter l'item choisi dans l'inventaire
	selectedItem := shopItems[choix-1]
	personnage.AddInventory(selectedItem)
	fmt.Println(selectedItem, "a été ajouté à votre inventaire !")
}
