package librairies

import "fmt"

// Ajoute un sort au personnage
func (c *Character) SpellBook(spell string) {
	// Vérifie si le sort existe déjà
	for _, s := range c.Skills {
		if s == spell {
			fmt.Printf("⚠️ %s connaît déjà le sort : %s\n", c.Name, spell)
			return
		}
	}

	// Ajoute le sort
	c.Skills = append(c.Skills, spell)
	fmt.Printf("✨ %s a appris un nouveau sort : %s\n", c.Name, spell)
}

// Affiche tous les sorts connus
func (c *Character) ShowSkills() {
	if len(c.Skills) == 0 {
		fmt.Println("❌ Aucun sort connu.")
		return
	}

	fmt.Println("\n📜 Liste des sorts connus :")
	for i, s := range c.Skills {
		fmt.Printf("%d) %s\n", i+1, s)
	}
}