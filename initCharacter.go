// Créez une fonction initCharacter permettant d’initialiser un personnage en
// utilisant la structure Character.
// Dans votre fonction Main, initialisez un personnage c1 à l’aide de la
// fonction initCharacter avec les valeurs suivantes :
// Nom : votre nom
// Classe: Elfe
// Niveau : 1
// Points de vie maximum : 100
// Points de vie actuels : 40
// Inventaire : 3 potions

package main 

func initCharacter(name string, classe string, level int, max_pv int, pv int, inventory [10]string) Character { 
	return Character{
		Nom:        nom,
		Classe:     classe,
		Niveau:     niveau,
		PVMax:      pvMax,
		PVActuels:  pvActuels,
		Inventaire: inventaire,
	}
}