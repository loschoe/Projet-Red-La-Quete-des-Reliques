package librairies 

import(
	"fmt"
	"github.com/fatih/color"
)

type Character struct {
	Name               string
	Classe             string
	Level              int
	Max_PV             int
	PV                 int
	Inventory          [10]string
	HasReceivedDiamond bool
	Rubis              int
	Attack             int
}

func InitCharacter(name string, classe string, level int, max_pv int, pv int, inventory [10]string) Character {
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
		Rubis:     15, // Commence avec 15 rubis
		Attack:    6, // Attaque de 6
	}
}

func DisplayInfo(c *Character) {
    fmt.Printf("\nName : %s\nClasse : %s\nLevel : %d\nPV : %d/%d\nInventory : %v\nRubis : %d\n",
        c.Name, c.Classe, c.Level, c.PV, c.Max_PV, c.Inventory, c.Rubis)
}

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

func (c *Character) RemoveItem(item string) {
	for idx, i := range c.Inventory {
		if i == item {
			c.Inventory[idx] = ""
			fmt.Println(item, "a été retiré de l'inventaire. \n")
			return
		}
	}
}

func (c *Character) CountItem(item string) int {
	count := 0
	for _, i := range c.Inventory {
		if i == item {
			count++
		}
	}
	return count
}

func (c *Character) RemoveItemAt(index int) {
	for j := index; j < len(c.Inventory)-1; j++ {
		c.Inventory[j] = c.Inventory[j+1]
	}
	c.Inventory[len(c.Inventory)-1] = ""
}

func (c *Character) IsInventoryFull() bool {
	count := 0
	for _, item := range c.Inventory {
		if item != "" && item != "..." {
			count++
		}
	}
	return count >= 10
}

func (personnage *Character) IsDead() {
    if personnage.PV <= 0 {
        color.HiRed("%s est mort ! ⚰️", personnage.Name)
        personnage.PV = personnage.Max_PV / 2
        color.Green("%s est ressuscité avec %d/%d PV ! ✨", personnage.Name, personnage.PV, personnage.Max_PV)
    }
}

// Utiliser le premier objet de l'inventaire
func (c *Character) UseItem() {
	switch c.Inventory[0] {
	case "Fairy":
		c.TakePot() // vient de potions.go
		c.Inventory[0] = "" // supprime l'objet
	case "Miasme":
		c.PoisonPot() // vient de potions.go
		c.Inventory[0] = ""
	default:
		fmt.Println("Aucun objet utilisable sélectionné !")
	}
}

// Utiliser un objet à un index choisi
func (c *Character) UseItemAt(index int) {
	if index < 0 || index >= len(c.Inventory) || c.Inventory[index] == "" || c.Inventory[index] == "..." {
		fmt.Println("Case invalide ou vide !")
		return
	}

	item := c.Inventory[index]
	switch item {
	case "Fairy":
		c.TakePot()
		c.Inventory[index] = ""
	case "Miasme":
		c.PoisonPot()
		c.Inventory[index] = ""
	default:
		fmt.Println("Cet objet ne peut pas être utilisé !")
	}
}
