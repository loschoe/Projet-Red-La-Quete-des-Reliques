// Ce fichier contient la fonction clear console pour un affichage dynamique 
// Le paquet de la librairie où sont stockées les fonctions 
package librairies

import(
	"os"
    "os/exec"
	"fmt"
	"github.com/fatih/color"
)

func ClearScreen() {
    var cmd *exec.Cmd
    if os.PathSeparator == '\\' {
        cmd = exec.Command("cmd", "/c", "cls") // Windows
    } else {
        cmd = exec.Command("clear") // Linux/Mac
    }
    cmd.Stdout = os.Stdout
    cmd.Run()
}

func Pause() {
    color.Red("\nPress Enter...")
    fmt.Scanln()
}
