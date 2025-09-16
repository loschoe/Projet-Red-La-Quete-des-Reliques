package librairies

import (
	"fmt"
	"github.com/fatih/color"
	"time"
)

// Monster spécifique pour le combat
type Monster struct {
	Name   string
	Max_PV int
	PV     int
	Attack int
}

// Initialisation d’un Bokoblin
func InitBokoblin(name string, max_PV int, pv int, attack int) Monster {
	if pv > max_PV {
		pv = max_PV
	}
	return Monster{
		Name:   name,
		Max_PV: max_PV,
		PV:     pv,
		Attack: attack,
	}
}

// Initialisation d’un Moblin
func InitMoblin(name string, max_PV int, pv int, attack int) Monster {
	if pv > max_PV {
		pv = max_PV
	}	
	return Monster{
		Name:   name,
		Max_PV: max_PV,
		PV:     pv,
		Attack: attack,
	}
}

// Initialisation d’un Lynel
func InitLynel(name string, max_PV int, pv int, attack int) Monster {
	if pv > max_PV {
		pv = max_PV
	}	
	return Monster{
		Name:   name,
		Max_PV: max_PV,	
		PV:     pv,
		Attack: attack,
	}
}	


// Initialisation d’un Krrrooçe
func InitKrrooçe(name string, max_PV int, pv int, attack int) Monster {
	if pv > max_PV {
		pv = max_PV
	}
	return Monster{
		Name:   name,
		Max_PV: max_PV,	
		PV:     pv,
		Attack: attack,
	}
}


// Attaque du Bokoblin (pattern)
func (m *Monster) GoblinPattern(player *Character, turn int) {
	damage := m.Attack
	if turn%3 == 0 {
		damage = m.Attack * 2
	}

	player.PV -= damage
	if player.PV < 0 {
		player.PV = 0
	}

	color.Red("\n%s inflige %d de dégâts à %s !\n", m.Name, damage, player.Name)
	color.Blue("PV de %s : %d/%d\n\n", player.Name, player.PV, player.Max_PV)
}

// Attaque du Moblin (pattern)
func (m *Monster) MoblinPattern(player *Character, turn int) {
	damage := m.Attack

	// Exemple : Moblin inflige plus de dégâts tous les 4 tours
	if turn%4 == 0 {
		damage = m.Attack * 2
	}

	player.PV -= damage
	if player.PV < 0 {
		player.PV = 0
	}

	color.Red("\n%s attaque %s et inflige %d de dégâts !\n", m.Name, player.Name, damage)
	color.Blue("PV de %s : %d/%d\n\n", player.Name, player.PV, player.Max_PV)
}

// Attaque du Lynel (pattern)
func (m *Monster) LynelPattern(player *Character, turn int) {
	damage := m.Attack

	// Lynel attaque fort tous les 5 tours
	if turn%5 == 0 {
		damage = m.Attack * 3
	} else if turn%2 == 0 {
		damage = m.Attack * 2
	}

	player.PV -= damage
	if player.PV < 0 {
		player.PV = 0
	}

	color.Red("\n%s fonce sur %s et inflige %d de dégâts !\n", m.Name, player.Name, damage)
	color.Blue("PV de %s : %d/%d\n\n", player.Name, player.PV, player.Max_PV)
}

// Attaque du Krrrooçe (pattern)
func (m *Monster) KrrrooçePattern(player *Character, turn int) {
	damage := m.Attack

	// Krrrooçe attaque aléatoirement plus fort
	if turn%6 == 0 {
		damage = m.Attack * 4
	} else if turn%3 == 0 {
		damage = m.Attack * 2
	}

	player.PV -= damage
	if player.PV < 0 {
		player.PV = 0
	}

	color.Red("\n%s frappe %s avec rage et inflige %d de dégâts !\n", m.Name, player.Name, damage)
	color.Blue("PV de %s : %d/%d\n\n", player.Name, player.PV, player.Max_PV)
}

	
// Lancer un combat générique pour Moblin, Lynel ou Krrrooçe
func StartFight(player *Character, monster *Monster, pattern func(*Character, int)) {
	turn := 1
	for player.PV > 0 && monster.PV > 0 {
		CharTurn(player, monster, turn)
		if monster.PV > 0 && player.PV > 0 {
			pattern(monster, player, turn)
		}
		turn++
	}

	if player.PV <= 0 {
		color.Red("\n%s a été vaincu 💀\n", player.Name)
		color.Yellow("Retour au menu principal...\n")
	} else if monster.PV <= 0 {
		color.Yellow("\n%s a été vaincu 🎉\n", monster.Name)
		player.Rubis += 50
		color.Green("%s reçoit 50 rubis !\n", player.Name)
		player.Level += 1
		color.Green("%s passe au niveau supérieur !\n", player.Name)
		color.Yellow("\nRetour au menu principal...\n")
	}
}

// Menu de sélection du monstre
func SelectMonsterFight(player *Character) {
	color.Red("\n============== Choix du combat ==============\n")
	color.Yellow("1. Combattre un Bokoblin (entraînement)\n")
	color.Yellow("2. Combattre un Moblin\n")
	color.Yellow("3. Combattre un Lynel\n")
	color.Yellow("4. Combattre un Krrrooçe\n")
	color.Yellow("5. Retour au menu principal\n")
	fmt.Print("\nVotre choix : ")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		TrainingFight(player)
	case 2:
		moblin := InitMoblin("Moblin", 60, 60, 8)
		StartFight(player, &moblin, (*Monster).MoblinPattern)
	case 3:
		lynel := InitLynel("Lynel", 120, 120, 15)
		StartFight(player, &lynel, (*Monster).LynelPattern)
	case 4:
		krrrooce := InitKrrooçe("Krrrooçe", 150, 150, 20)
		StartFight(player, &krrrooce, (*Monster).KrrrooçePattern)
	case 5:
		fmt.Println("Retour au menu principal...")
	default:
		fmt.Println("Choix invalide !")
	}
}

// Tour du joueur
func CharTurn(player *Character, monster *Monster, turn int) {
	shopArt := `
   _____                _           _   
  / ____|              | |         | |  
 | |     ___  _ __ ___ | |__   __ _| |_ 
 | |    / _ \| '_ ' _ \| '_ \ / _' | __|
 | |___| (_) | | | | | | |_) | (_| | |_ 
  \_____\___/|_| |_| |_|_.__/ \__,_|\__|
`
	color.Red("%s\n", shopArt)

	color.Blue("\n============== Tour %d==============\n", turn)

	time.Sleep(1 * time.Second)
	color.Green("Adversaire : %s | PV : %d/%d\n", monster.Name, monster.PV, monster.Max_PV)
	color.Green("%s | PV : %d/%d\n\n", player.Name, player.PV, player.Max_PV)

	// Nouveau menu avec 4 choix
	color.Red("1. Attaquer")
	color.Yellow("2. Inventaire / Utiliser objet")
	color.Cyan("3. Choisir le boss à affronter")
	color.Magenta("4. Fuir (retour menu)")
	fmt.Print("\nVotre choix : ")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1: // Attaque
		damage := player.Attack
		monster.PV -= damage
		if monster.PV < 0 {
			monster.PV = 0
		}
		color.Red("\n%s attaque et inflige %d dégâts à %s !\n", player.Name, damage, monster.Name)
		time.Sleep(1 * time.Second)
		color.Green("PV restants de %s : %d/%d\n\n", monster.Name, monster.PV, monster.Max_PV)

	case 2: // Inventaire
		player.AccessInventory()
		fmt.Print("Choisissez un objet à utiliser : ")
		var itemChoice int
		fmt.Scanln(&itemChoice)

		if itemChoice < 1 || itemChoice > len(player.Inventory) {
			fmt.Println("Choix invalide !")
			return
		}

		item := player.Inventory[itemChoice-1]

		switch item {
		case "Fairy":
			player.TakePot()
			player.RemoveItemAt(itemChoice - 1)

		case "Miasme":
			var choix int
			fmt.Println("Que voulez-vous faire avec le poison ? (1 = boire, 2 = lancer sur le monstre) :")
			fmt.Scanln(&choix)

			if choix == 1 {
				player.Poisonbottle()
			} else if choix == 2 {
				player.PoisonPot(monster)
			} else {
				fmt.Println("Choix invalide, aucune action effectuée.")
			}

		default:
			fmt.Println("Cet objet ne peut pas être utilisé !")
		}

	case 3: // Choisir le boss
		SelectMonsterFight(player)

	case 4: // Fuite
		fmt.Println("Vous prenez la fuite... retour au menu principal.")
		player.PV = 0

	default:
		fmt.Println("Choix invalide, vous perdez votre tour !")
	}
}


// Combat d’entraînement
func TrainingFight(player *Character) {
	bokoblin := InitBokoblin("Bokoblin", 40, 40, 5)

	turn := 1
	for player.PV > 0 && bokoblin.PV > 0 {
		CharTurn(player, &bokoblin, turn)

		if bokoblin.PV > 0 && player.PV > 0 {
			bokoblin.GoblinPattern(player, turn)
		}
		turn++
	}

	if player.PV <= 0 {
		color.Red("\n%s a été vaincu 💀\n", player.Name)
		color.Yellow("Retour au menu principal...\n")
	} else if bokoblin.PV <= 0 {
		color.Yellow("\n%s a été vaincu 🎉\n", bokoblin.Name)
		player.Rubis += 25
		color.Green("%s reçoit 25 rubis !\n", player.Name)
		player.Level += 1
		color.Green("%s passe au niveau supérieur !\n", player.Name)
		color.Yellow("\nRetour au menu principal...\n")
	}
}