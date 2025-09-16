package librairies

import (
	"fmt"
	"github.com/fatih/color"
	"time"
)

// ------------------- Structures -------------------

type Monster struct {
	Name   string
	Max_PV int
	PV     int
	Attack int
}

// ------------------- Initialisations -------------------

func InitBokoblin(name string, max_PV int, pv int, attack int) Monster {
	if pv > max_PV {
		pv = max_PV
	}
	return Monster{Name: name, Max_PV: max_PV, PV: pv, Attack: attack}
}

func InitMoblin(name string, max_PV int, pv int, attack int) Monster {
	if pv > max_PV {
		pv = max_PV
	}
	return Monster{Name: name, Max_PV: max_PV, PV: pv, Attack: attack}
}

func InitLynel(name string, max_PV int, pv int, attack int) Monster {
	if pv > max_PV {
		pv = max_PV
	}
	return Monster{Name: name, Max_PV: max_PV, PV: pv, Attack: attack}
}

func InitKrrooçe(name string, max_PV int, pv int, attack int) Monster {
	if pv > max_PV {
		pv = max_PV
	}
	return Monster{Name: name, Max_PV: max_PV, PV: pv, Attack: attack}
}

// ------------------- Patterns -------------------

func (m *Monster) GoblinPattern(player *Character, turn int) {
	damage := m.Attack
	if turn%3 == 0 {
		damage *= 2
	}
	player.PV -= damage
	if player.PV < 0 {
		player.PV = 0
	}
	color.Red("\n%s inflige %d de dégâts à %s !\n", m.Name, damage, player.Name)
	color.Blue("PV de %s : %d/%d\n\n", player.Name, player.PV, player.Max_PV)
}

func (m *Monster) MoblinPattern(player *Character, turn int) {
	damage := m.Attack
	if turn%4 == 0 {
		damage *= 2
	}
	player.PV -= damage
	if player.PV < 0 {
		player.PV = 0
	}
	color.Red("\n%s attaque %s et inflige %d de dégâts !\n", m.Name, player.Name, damage)
	color.Blue("PV de %s : %d/%d\n\n", player.Name, player.PV, player.Max_PV)
}

func (m *Monster) LynelPattern(player *Character, turn int) {
	damage := m.Attack
	if turn%5 == 0 {
		damage *= 3
	} else if turn%2 == 0 {
		damage *= 2
	}
	player.PV -= damage
	if player.PV < 0 {
		player.PV = 0
	}
	color.Red("\n%s fonce sur %s et inflige %d de dégâts !\n", m.Name, player.Name, damage)
	color.Blue("PV de %s : %d/%d\n\n", player.Name, player.PV, player.Max_PV)
}

func (m *Monster) KrrrooçePattern(player *Character, turn int) {
	damage := m.Attack
	if turn%10 == 0 {
		damage *= 4
	} 
	player.PV -= damage
	if player.PV < 0 {
		player.PV = 0
	}
	color.Red("\n%s frappe %s avec rage et inflige %d de dégâts !\n", m.Name, player.Name, damage)
	color.Blue("PV de %s : %d/%d\n\n", player.Name, player.PV, player.Max_PV)
}

// ------------------- Tour du joueur -------------------

func CombatTurn(player *Character, monster *Monster, turn int) {
	shopArt := `
   _____                _           _   
  / ____|              | |         | |  
 | |     ___  _ __ ___ | |__   __ _| |_ 
 | |    / _ \| '_ ' _ \| '_ \ / _' | __|
 | |___| (_) | | | | | | |_) | (_| | |_ 
  \_____\___/|_| |_| |_|_.__/ \__,_|\__|
`
	color.Red("%s\n", shopArt)
	color.Blue("\n============== Tour %d ==============\n", turn)
	time.Sleep(1 * time.Second)
	color.Green("Adversaire : %s | PV : %d/%d\n", monster.Name, monster.PV, monster.Max_PV)
	color.Green("%s | PV : %d/%d\n\n", player.Name, player.PV, player.Max_PV)

	color.Red("1. Attaque Coup de poing")
	color.Red("2. Epée tranchante")
	color.Red("3. Pluie de flèches")
	color.Yellow("4. Inventaire / Utiliser objet")
	color.Cyan("5. Fuir (retour menu)")
	fmt.Print("\nVotre choix : ")

	var choice int
	fmt.Scanln(&choice)

	switch choice {

	case 1:
		damage := player.Attack
		monster.PV -= damage
		if monster.PV < 0 {
			monster.PV = 0
		}
		color.Red("\n%s attaque et inflige %d dégâts à %s !\n", player.Name, damage, monster.Name)
		time.Sleep(1 * time.Second)
		color.Green("PV restants de %s : %d/%d\n\n", monster.Name, monster.PV, monster.Max_PV)

	case 2:
		fmt.Printf("\n%s utilise sa Master Sword !\n", player.Name)
		player.UseMasterSword(monster)
		time.Sleep(1 * time.Second)

	case 3:
		fmt.Printf("\n%s utilise son Arc !\n", player.Name)
		player.UseBow(monster)
		time.Sleep(1 * time.Second)

	case 4:
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

	case 5:
		fmt.Println("Vous prenez la fuite... retour au menu principal.")
		player.PV = 0

	default:
		fmt.Println("Choix invalide, vous perdez votre tour !")
	}
}

func CharTurn(player *Character, monster *Monster, turn int) {
	CombatTurn(player, monster, turn)
}

// ------------------- Menu Combat -------------------
func CombatMenu(player *Character) {
	color.Cyan("+---------------------------------+")
    color.Cyan("|        Choix du combat          |")
	color.Cyan("+---------------------------------+")
    color.Yellow("|⚔️  Duel d'entraînement (Bokoblin)|")
    color.Blue("|💀 Combat Boss (Moblin)          |")
    color.Green("|🐴 Combat Boss (Lynel)           |")
    color.Red("|🍔 Combat Boss (Krrooçe)         |")
	color.Cyan("+---------------------------------+")
    fmt.Print("\nVotre choix : ")

	var choice int
	fmt.Scanln(&choice)

    switch choice {
    case 1:
        TrainingFight(player)
    case 2:
        StartFight(player, InitMoblin("Moblin", 70, 70, 20), (*Monster).MoblinPattern)
    case 3:
        StartFight(player, InitLynel("Lynel", 150, 150, 50), (*Monster).LynelPattern)
    case 4:
        StartFight(player, InitKrrooçe("Krrooçe", 450, 450, 150), (*Monster).KrrrooçePattern)
    case 5:
        fmt.Println("Retour au menu principal")
    default:
        fmt.Println("Choix invalide !")
    }
}


// ------------------- Combat Bokoblin -------------------

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

// ------------------- Combat Boss -------------------

func StartFight(player *Character, monster Monster, pattern func(*Monster, *Character, int)) {
	turn := 1
	for player.PV > 0 && monster.PV > 0 {
		CharTurn(player, &monster, turn)
		if monster.PV > 0 && player.PV > 0 {
			pattern(&monster, player, turn)
		}
		turn++
	}
	if player.PV <= 0 {
		color.Red("\n%s a été vaincu 💀\n", player.Name)
	} else if monster.PV <= 0 {
		color.Yellow("\n%s a été vaincu 🎉\n", monster.Name)
		switch monster.Name {
		case "Moblin":
			// Si c’est un Moblin
			player.Rubis += 30
			player.Level += 2
			color.Green("%s reçoit 30 rubis et passe de 2 niveaux !\n", player.Name)
		case "Lynel":
			drops := []string{"Diamant", "Tissu Royal"}
			for _, item := range drops {
				for i := 0; i < len(player.Inventory); i++ {
					if player.Inventory[i] == "" || player.Inventory[i] == "..." {
						player.Inventory[i] = item
						color.Green("%s reçoit : %s\n", player.Name, item)
						break
					}
				}
			}
		case "Krrooçe":
			// Si c’est Krrrooçe
			color.HiRed("\n%s a été vaincu ! FIN DU JEU 🎉\n", monster.Name)
			player.GameOver = true
		default:
			// Si c’est un autre monstre
			player.Rubis += 50
			player.Level += 1
			color.Green("%s reçoit 50 rubis et passe un niveau !\n", player.Name)
		}
	}
}