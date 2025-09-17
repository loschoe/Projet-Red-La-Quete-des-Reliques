// Ce fichier contient toutes les fonctions nécéssaires au fonctionnement des items qui sont présent dans le jeu. 
// Le paquet de la librairie où sont stockées les fonctions. 

package librairies

import (
	"github.com/fatih/color" // Couleurs 
	"fmt" 				  // Pour les prints
)

//=================== EPEE ==================================
// Nombre d'utilisation de l'épée
var MasterSwordUses = 100

func (personnage *Character) UseMasterSword(monster *Monster) {
	for i, item := range personnage.Inventory {
		if item == "Master Sword" {
			println(personnage.Name, "utilise l'épée purificatrice !\n")

			// Inflige 50 points de dégâts 
			monster.PV -= 50
			if monster.PV < 0 {
				monster.PV = 0
			}
			color.Red("%s inflige 50 points de dégâts ! PV restant du monstre : %d / %d\n",
				personnage.Name, monster.PV, monster.Max_PV)

			// Réduire le compteur d'usage 
			if MasterSwordUses > 0 {
				MasterSwordUses--
				color.HiBlack("Usages restants de la Master Sword : %d\n", MasterSwordUses)
			} else {
				color.HiBlack("La Master Sword est usée !\n")
				// Retirer l'épée si usages = 0
				personnage.RemoveItemAt(i)
			}

			return
		}
	}
	color.Red("Vous n'avez pas l'épée dans votre inventaire.\n")
}

//=================== ARC ==========================
var BowUses = 4 // Compteur d'usage 

func (personnage *Character) UseBow(monster *Monster) {
	// Vérifie si le personnage a un arc
	arcIndex := -1
	for i, item := range personnage.Inventory {
		if item == "Bow" {
			arcIndex = i
			break
		}
	}

	if arcIndex == -1 {
		color.Red("Vous n'avez pas l'arc dans votre inventaire.\n")
		return
	}

	// Vérifie si le personnage a une flèche
	arrowIndex := -1
	for i, item := range personnage.Inventory {
		if item == "Arrow" {
			arrowIndex = i
			break
		}
	}

	if arrowIndex == -1 {
		color.Red("Vous n'avez plus de flèches.\n")
		return
	}

	println(personnage.Name, "utilise un arc !\n")

	// Inflige 100 points de dégâts
	monster.PV -= 100
	if monster.PV < 0 {
		monster.PV = 0
	}
	color.Red("%s inflige 100 points de dégâts ! PV restant du monstre : %d / %d\n",
		personnage.Name, monster.PV, monster.Max_PV)

	// Réduire le compteur d'usage de l'arc
	if BowUses > 0 {
		BowUses--
		color.HiBlack("Usages restants de l'arc : %d\n", BowUses)
	} else {
		color.HiBlack("L'arc est usé !\n")
		// Retirer l'arc si usages = 0
		personnage.RemoveItemAt(arcIndex)
	}

	// Retirer la flèche utilisée
	personnage.RemoveItemAt(arrowIndex)
}

//=================== Sort ==========================
func (player *Character) UseFireBall(monster *Monster) {
    if player.FireBallUsed {
        fmt.Println("❌ Vous avez déjà utilisé Zelda Book dans ce combat !")
        return
    }

    // Inflige 70 dégâts au monstre
    damage := 70
    monster.PV -= damage
    if monster.PV < 0 {
        monster.PV = 0
    }

    // Restaure 15 PV au joueur
    heal := 15
    player.PV += heal
    if player.PV > player.Max_PV {
        player.PV = player.Max_PV
    }

    // Affichage
    color.Red("\n%s appelle la princesse Zelda à son secours et Zelda inflige %d dégâts à %s !\n", player.Name, damage, monster.Name)
    color.Green("%s récupère %d PV ! PV actuels : %d/%d\n\n", player.Name, heal, player.PV, player.Max_PV)
    color.Green("PV restants de %s : %d/%d\n\n", monster.Name, monster.PV, monster.Max_PV)

    player.FireBallUsed = true // ← marque Fire Ball comme utilisée
}

// ================== Equipement ================================
// Vérifie si un objet est équipé
func (c *Character) HasEquipment(item string) bool {
    for _, i := range c.Equipment {
        if i == item {
            return true
        }
    }
    return false
}

// Applique les bonus des équipements au personnage
func (c *Character) ApplyEquipmentBonus() {
    if c.EquipmentApplied == nil {
        c.EquipmentApplied = make(map[string]bool)
    }

    bonuses := map[string]int{
        "Casque de garde": 75,
        "Tunique royale":  300,
        "Bottes":         25,
    }

    for item, bonus := range bonuses {
        if c.HasEquipment(item) && !c.EquipmentApplied[item] {
            c.PV += bonus
            color.Green("Votre vie augmente de %d PV grâce à %s !\n", bonus, item)
            c.EquipmentApplied[item] = true
        }
    }
}