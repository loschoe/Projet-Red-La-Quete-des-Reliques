// Ce fichier contient toutes les fonctions nécéssaires au personnage et à son fonctionnement 
// Le paquet de la librairie où sont stockées les fonctions 

package librairies

import (
	"fmt"              // Pour les prints
	"github.com/fatih/color" // Pour afficher en couleurs dans la console
)

// Structure joueur
type Character struct {
	Name               string
	Classe             string
	Level              int
	Max_PV             int
	PV                 int
	Inventory          []string  // Stocker les items dans un slice
	InventoryCapacity  int       // Capacité de l'inventaire
	InventoryUpgrades  int       // 3 améliorations possibles
	Equipment          [3]string // Stocker les tenues dans un array de 3 cases
	HasReceivedDiamond bool      // Cadeau du marchand lors du 1er passage
	Rubis              int       // Argent du jeu
	Attack             int       // Les points d'attaque du joueur
	GameOver           bool      // Détection de la fin du jeu ou d'une mort
}

// ------------- INITIALISATION -------------------------

// Initialisation d'un personnage
func InitCharacter(name string, classe string, level int, max_pv int, pv int, inventory [10]string, equipment [3]string) Character {
	if pv > max_pv {
		pv = max_pv
	}

	// Initialisation de l'inventaire
	baseInventory := make([]string, 10)
	for i := 0; i < len(inventory) && i < 10; i++ {
		baseInventory[i] = inventory[i]
	}

	// Initialisation de l'équipement
	var baseEquipment [3]string
	for i := 0; i < len(equipment) && i < 3; i++ {
		baseEquipment[i] = equipment[i]
	}

	return Character{
		Name:              name,
		Classe:            classe,
		Level:             level,
		Max_PV:            max_pv,
		PV:                pv,
		Inventory:         baseInventory,
		InventoryCapacity: 10,
		InventoryUpgrades: 0,
		Equipment:         baseEquipment,
		Rubis:             900,
		Attack:            6,
		GameOver:          false,
	}
}

// Affiche les infos du personnage
func DisplayInfo(c *Character) {
	fmt.Printf("\nName : %s\nClasse : %s\nLevel : %d\nPV : %d/%d \nAttack : %v \nInventory : %v\nRubis : %d\n",
		c.Name, c.Classe, c.Level, c.PV, c.Max_PV, c.Attack, c.Inventory, c.Rubis)
}

// ------------- INVENTAIRE -------------------------

// Accéder à l'inventaire
func (c *Character) AccessInventory() {
	fmt.Println("\nInventaire du personnage :")
	vide := true
	for i, item := range c.Inventory {
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
	fmt.Println("Rubis disponibles :", c.Rubis)
	fmt.Println()
}

// Ajouter un objet dans l'inventaire
func (c *Character) AddInventory(item string) {
	if c.IsInventoryFull() {
		fmt.Println("Inventaire plein ! Impossible d'ajouter", item)
		return
	}
	for i := 0; i < len(c.Inventory); i++ {
		if c.Inventory[i] == "" || c.Inventory[i] == "..." {
			c.Inventory[i] = item
			fmt.Println(item, "a été ajouté à l'inventaire.")
			return
		}
	}
}

// Augmenter la taille de l'inventaire (3x max, +5 cases à chaque fois)
func (c *Character) UpgradeInventorySlot() {
	if c.InventoryUpgrades >= 3 {
		fmt.Println("❌ Vous avez déjà atteint la limite d’augmentations (3).")
		return
	}

	c.InventoryCapacity += 5
	c.InventoryUpgrades++

	newInventory := make([]string, c.InventoryCapacity)
	copy(newInventory, c.Inventory)
	c.Inventory = newInventory

	fmt.Printf("✅ Votre inventaire a été agrandi ! Nouvelle capacité : %d slots\n", c.InventoryCapacity)
}

// Retirer un objet par son nom
func (c *Character) RemoveItem(item string) {
	for idx, i := range c.Inventory {
		if i == item {
			c.Inventory[idx] = ""
			fmt.Println(item, "a été retiré de l'inventaire.\n")
			return
		}
	}
}

// Retirer un objet à un index
func (c *Character) RemoveItemAt(index int) {
	for j := index; j < len(c.Inventory)-1; j++ {
		c.Inventory[j] = c.Inventory[j+1]
	}
	c.Inventory[len(c.Inventory)-1] = ""
}

// Compter combien de fois un objet apparaît
func (c *Character) CountItem(item string) int {
	count := 0
	for _, i := range c.Inventory {
		if i == item {
			count++
		}
	}
	return count
}

// Vérifier si l'inventaire est plein
func (c *Character) IsInventoryFull() bool {
	count := 0
	for _, item := range c.Inventory {
		if item != "" && item != "..." {
			count++
		}
	}
	return count >= c.InventoryCapacity
}

// ------------- EQUIPEMENT -------------------------

// Accéder à l'équipement
func (c *Character) AccessEquipment() {
	fmt.Println("\nÉquipement du personnage :")
	vide := true
	for i, item := range c.Equipment {
		if item == "" || item == "..." {
			fmt.Printf("%d. [vide]\n", i+1)
		} else {
			fmt.Printf("%d. %s\n", i+1, item)
			vide = false
		}
	}
	if vide {
		fmt.Println("Équipement vide.")
	}
	fmt.Println()
}

// Ajouter un équipement
func (c *Character) AddEquipment(item string) {
	for i := 0; i < len(c.Equipment); i++ {
		if c.Equipment[i] == "" || c.Equipment[i] == "..." {
			c.Equipment[i] = item
			fmt.Println(item, "a été ajouté à l'équipement.")
			return
		}
	}
	fmt.Println("Impossible d'ajouter l'équipement :", item, "(tous les slots sont pleins)")
}

// ------------- ETAT DU PERSONNAGE -------------------------

// Vérifier si le personnage est mort
func (c *Character) IsDead() {
	if c.PV <= 0 {
		color.HiRed("%s est mort ! ⚰️\n", c.Name)
		c.PV = c.Max_PV / 2
		color.Green("%s est ressuscité avec %d/%d PV ! ✨\n", c.Name, c.PV, c.Max_PV)
	}
}
