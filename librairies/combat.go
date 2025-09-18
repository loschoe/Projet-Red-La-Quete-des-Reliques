// Ce fichier contient toutes les fonctions nécéssaires au combat et à son fonctionnement 
// Le paquet de la librairie où sont stockées les fonctions 

// Les combats se terminent soient par la mort soit par la fuite soit par la victoire du joueur !
package librairies

import (
	"fmt" 						//Certains affichages 
	"github.com/fatih/color"    //Les couleurs en console 
	"time"                      //Dans certaines attaques il nous faut une source de temps 
)

// ------------------- Structures -------------------

type Monster struct {
	Name   string
	Max_PV int
	PV     int
	Attack int
}

// ------------------- Initialisations -------------------

// 1er monstre Bokoblin
func InitBokoblin(name string, max_PV int, pv int, attack int) Monster {
	if pv > max_PV {
		pv = max_PV
	}
	return Monster{Name: name, Max_PV: max_PV, PV: pv, Attack: attack}
}

// 2e monstre Moblin
func InitMoblin(name string, max_PV int, pv int, attack int) Monster {
	if pv > max_PV {
		pv = max_PV
	}
	return Monster{Name: name, Max_PV: max_PV, PV: pv, Attack: attack}
}

// 3e monstre Lynel 
func InitLynel(name string, max_PV int, pv int, attack int) Monster {
	if pv > max_PV {
		pv = max_PV
	}
	return Monster{Name: name, Max_PV: max_PV, PV: pv, Attack: attack}
}

// 4e monstre BOSS FINAL (le nom est bizarre, c'est normal c'est un délire)
func InitKrrooçe(name string, max_PV int, pv int, attack int) Monster {
	if pv > max_PV {
		pv = max_PV
	}
	return Monster{Name: name, Max_PV: max_PV, PV: pv, Attack: attack}
}

// ------------------- Patterns de combats -------------------

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

	fmt.Println("\n=== Attaques disponibles ===")

	// Affiche toutes les compétences de la classe choisie
	for i, skill := range player.Skills {
    	color.Red("%d. Attaque %s", i+1, skill)
	}

	// Master Sword → débloque "Épée tranchante"
	if player.HasItem("Master Sword") {
    	color.Red("2. Épée tranchante")
	}

	// Arc → débloque "Pluie de flèches"
	if player.HasItem("Bow") {
    	color.Red("3. Pluie de flèches")
	}

	// Zelda Book → débloque "Boule de feu" 
	if player.HasItem("Zelda Book") && !player.FireBallUsed {
    color.Red("4. Zelda Book")
	}

	// Inventaire
	color.Yellow("5. Inventaire / Utiliser objet")

	// Fuite
	color.Cyan("6. Fuir (retour menu)")

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
		player.UseMasterSword(monster)
		time.Sleep(1 * time.Second)

	case 3:
		player.UseBow(monster)
		time.Sleep(1 * time.Second)

	case 4:
    if !player.FireBallUsed {
        player.UseFireBall(monster)
        player.FireBallUsed = true 
    } else {
        fmt.Println("Vous avez déjà utilisé Fire Ball ce combat !")
    }
    time.Sleep(1 * time.Second)


	case 5:
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

	case 6:
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
	player.FireBallUsed = false
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
		player.Rubis += 10
		color.Green("%s reçoit 10 rubis !\n", player.Name)
		player.Level += 1
		color.Green("%s passe au niveau supérieur !\n", player.Name)
		//  Augmentation de l'attaque si Bokoblin vaincu
		player.Attack += 4
		color.HiMagenta("💥 Votre attaque augmente de 4 ! Nouvelle attaque : %d\n", player.Attack)
		color.Yellow("\nRetour au menu principal...\n")
	}
}

// ------------------- Combat Boss -------------------

func StartFight(player *Character, monster Monster, pattern func(*Monster, *Character, int)) {
	player.FireBallUsed = false
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
			player.FireBallUsed = false
			// Si c’est un Moblin
			player.Rubis += 35
			player.Level += 2
			color.Green("%s reçoit 35 rubis et passe de 2 niveaux !\n", player.Name)
			// Si Moblin est mort, augmenter l'attaque de 5
			player.Attack += 5
			color.HiMagenta("💥 Votre attaque augmente de 5 ! Nouvelle attaque : %d\n", player.Attack)
		case "Lynel":
			drops := []string{"Diamant", "Tissu royal"}
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