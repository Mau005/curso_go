package capituloUno

import "fmt"

func SalidaStandar() {
	var texto string = "Yo soy la asignada"
	fmt.Println("Salida con un salto de linea es ln")
	fmt.Print("Esta es sin salto de linea")
	fmt.Printf("\n dar un formato a una variable como: |%v |esta es la variable", texto)

}
