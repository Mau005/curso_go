package capituloUno

import (
	"fmt"
)

func BuclesIteradores() {

	var texto string = "Bienvenido al mundo de los iteradores"

	for i := 0; i < 37; i++ {
		fmt.Printf("%c", texto[i])
	}
}
