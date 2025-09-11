package main   // toujours en haut

import (       // tous les imports doivent être ici
    "github.com/fatih/color"
)

func main() {  // tout code à exécuter doit être ici
    color.Red("Le rouge est la couleur de la passion et de l'énergie.")
    color.New(color.FgHiRed, color.Bold).Println("L'orange évoque la chaleur et la créativité.")
    color.Yellow("Le jaune illumine le monde avec optimisme et joie.")
    color.Green("Le vert symbolise la nature et l'harmonie.")
    color.Blue("Le bleu inspire la sérénité et la confiance.")
    color.Magenta("L'indigo nous invite à la réflexion et à l'intuition.")
    color.Cyan("Le violet incarne la sagesse et la spiritualité.")
}
