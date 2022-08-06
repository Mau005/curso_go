package main

import "fmt"

func main() {
	var condicion bool = true
	var opcion string = "no"

	for condicion == true {

		fmt.Println("Indicame si  para salir del bucle")
		fmt.Scan(&opcion)

		if opcion == "si" {
			condicion = false
		} else {
			fmt.Println("No se ha podido encontrar el filtro para estos caracrtreres")
		}
	}
	fmt.Println("Estoy fuera del bucle")
}
