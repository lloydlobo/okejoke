// [[file:../README.org::*okejoke/logo/logo.go][okejoke/logo/logo.go:1]]
package logo

// it works
import (
	"fmt"

	"github.com/fatih/color"
)

func PrintLogo() {
	okejoke := `
        _         _       _
   ___ | | _____ (_) ___ | | _____
  / _ \| |/ / _ \| |/ _ \| |/ / _ \
 | (_) |   <  __/| | (_) |   <  __/
  \___/|_|\_\___|/ |\___/|_|\_\___|
               |__/
    `
	color.Set(color.FgRed)
	fmt.Printf("%v\n", okejoke)
	color.Unset()
}

// okejoke/logo/logo.go:1 ends here
