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

	fmt.Printf("\n%s inflige %d de dégâts à %s !\n", m.Name, damage, player.Name)
	fmt.Printf("PV de %s : %d/%d\n\n", player.Name, player.PV, player.Max_PV)
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

	// Affichage du tour juste après l'ASCII
	color.Blue("\n===== Tour %d =====\n", turn)

	var choice int
	color.Green("------ Tour du joueur ------\n")
	color.Green("\nAdversaire : %s | PV : %d/%d\n", monster.Name, monster.PV, monster.Max_PV)
	color.Red("1. Attaquer\n")
	color.Yellow("2. Inventaire / Utiliser objet\n")
	color.Cyan("3. Fuir (retour menu)\n")
	fmt.Print("\nVotre choix : ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		damage := player.Attack
		monster.PV -= damage
		if monster.PV < 0 {
			monster.PV = 0
		}
		fmt.Printf("\n%s utilise Attaque coup de poing et inflige %d dégâts à %s !\n", player.Name, damage, monster.Name)
		time.Sleep(1 * time.Second)
		fmt.Printf("PV restants de %s : %d/%d\n\n", monster.Name, monster.PV, monster.Max_PV)
	case 2:
		player.AccessInventory()
		fmt.Print("Choisissez un objet à utiliser : ")
		var itemChoice int
		fmt.Scan(&itemChoice)
		player.UseItemAt(itemChoice - 1)
	case 3:
		color.Red("\nVous prenez la fuite... retour au menu principal.\n")
		player.PV = 0 // Force la sortie du combat
	default:
		fmt.Printf("\nChoix invalide, vous perdez votre tour !\n\n")
	}
}


// Combat d’entraînement
func TrainingFight(player *Character) {
	bokoblin := InitBokoblin("Bokoblin", 40, 40, 5)

	turn := 1
	for player.PV > 0 && bokoblin.PV > 0 {

		// Appel du tour du joueur avec l'affichage du tour après l'ASCII
		CharTurn(player, &bokoblin, turn)

		// Tour du monstre
		if bokoblin.PV > 0 && player.PV > 0 {
			bokoblin.GoblinPattern(player, turn)
		}

		turn++
	}

	// END FIGHT
	if player.PV <= 0 {
		color.Red("\n%s a été vaincu 💀\n", player.Name)
		color.Yellow("Retour au menu principal...\n")
	} else if bokoblin.PV <= 0 {
		color.Yellow("\n%s a été vaincu 🎉\n", bokoblin.Name)

		// Récompenses
		player.Rubis += 25
		color.Green("%s reçoit 25 rubis !\n", player.Name)
		player.Level += 1
		color.Green("%s passe au niveau supérieur !\n", player.Name)

		color.Yellow("\nRetour au menu principal...\n")
	}
}
