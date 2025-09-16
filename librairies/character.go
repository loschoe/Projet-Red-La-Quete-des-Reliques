// Ce fichier contient toutes les fonctions nécéssaires au personnage et à son fonctionnement 
// Le paquet de la librairie où sont stockées les fonctions 
package librairies

import (
	"fmt" 						// Pour les prints 
	"github.com/fatih/color" 	//Pour afficher en couleurs dans la console 
)

// Structure joueur
type Character struct {
	Name               string
	Classe             string
	Level              int
	Max_PV             int
	PV                 int
	Inventory          []string // Stocker les items dans un slice 
	InventoryCapacity  int      // Modifier la capacité de l'inventaire     
	InventoryUpgrades  int      // 3 améliorations possibles     
	HasReceivedDiamond bool     // Cadeau du marchand lors du 1er passage 
	Rubis              int      // Argent du jeu 
	Attack             int      // Les points d'attaque du joueur 
	GameOver		   bool     // Détection de la fin du jeu ou d'une mort 
}

// Initialisation d'un personnage
func InitCharacter(name string, classe string, level int, max_pv int, pv int, inventory [10]string) Character {
	if pv > max_pv {
		pv = max_pv
	}

	baseInventory := make([]string, 10) // 10 slots vides au départ
    for i := 0; i < len(inventory) && i < 10; i++ {
        baseInventory[i] = inventory[i]
    }

	return Character{				// Les variables de la structure sont reprises avec les valeurs 
		Name:      name,
		Classe:    classe,
		Level:     level,
		Max_PV:    max_pv,
		PV:        pv,
		Inventory:     baseInventory,
		InventoryCapacity: 10,
		InventoryUpgrades: 0,
		Rubis:     900,
		Attack:    6,
		GameOver : false,
	}
}

// Affiche les infos du personnage
func DisplayInfo(c *Character) {
	fmt.Printf("\nName : %s\nClasse : %s\nLevel : %d\nPV : %d/%d \nAttack : %v \nInventory : %v\nRubis : %d\n",
		c.Name, c.Classe, c.Level, c.PV, c.Max_PV, c.Attack, c.Inventory, c.Rubis)
}

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

// Augmente la taille de l'inventaire 3x (+ 5 cases)
func (c *Character) UpgradeInventorySlot() {
	if c.InventoryUpgrades >= 3 {
		fmt.Println("❌ Vous avez déjà atteint la limite d’augmentations (3).")
		return
	}

	c.InventoryCapacity += 5
	c.InventoryUpgrades++

	// recréer un nouvel inventaire avec p5 cases en plus 
	newInventory := make([]string, c.InventoryCapacity)
	copy(newInventory, c.Inventory)
	c.Inventory = newInventory

	fmt.Printf("✅ Votre inventaire a été agrandi ! Nouvelle capacité : %d slots\n", c.InventoryCapacity)
}


// Retirer un objet de l'inventaire par grâce au nom de l'item
func (c *Character) RemoveItem(item string) {
	for idx, i := range c.Inventory {
		if i == item {
			c.Inventory[idx] = ""
			fmt.Println(item, "a été retiré de l'inventaire. \n")
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

// Compter un objet
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


// Vérifier si le personnage est mort
func (c *Character) IsDead() {
	if c.PV <= 0 {
		color.HiRed("%s est mort ! ⚰️\n", c.Name)
		c.PV = c.Max_PV / 2
		color.Green("%s est ressuscité avec %d/%d PV ! ✨\n", c.Name, c.PV, c.Max_PV)
	}
}


// Utiliser un item en combat 
func (c *Character) UseItemAt(index int, monster *Monster) {
	if index < 0 || index >= len(c.Inventory) || c.Inventory[index] == "" || c.Inventory[index] == "..." {
		fmt.Println("Case invalide ou vide !")
		return
	}

	item := c.Inventory[index]

	switch item {
	case "Fairy":	// Une fée (potion de soin)
		c.TakePot()
		c.Inventory[index] = ""

	case "Miasme": // Un miasme (potion de poison)
		var choix int
		fmt.Println("Que voulez-vous faire avec le poison ? (1 = boire, 2 = lancer sur le monstre) :")
		fmt.Scan(&choix)

		if choix == 1 {
			c.Poisonbottle() 			// dégâts sur le joueur
		} else if choix == 2 {
			if monster != nil {
				c.PoisonPot(monster) 	// dégâts sur le monstre
			} else {
				fmt.Println("Aucun monstre cible pour Miasme !")
			}
		} else {
			fmt.Println("Choix invalide, aucune action effectuée.")
		}
	default:
		fmt.Println("Cet objet ne peut pas être utilisé !")
	}
}
