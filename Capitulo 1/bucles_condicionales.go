package main

import "fmt"

func main() {
	var condicion bool = true

	for condicion {
		var menu rune
		fmt.Println("Desea salir del bucle s/n")

		fmt.Scanf("%c\n", &menu)

		if menu == 's' {
			condicion = false
		} else if menu == 'n' {
			fmt.Println("Desea continuar con el bucle")
			continue
		} else {
			fmt.Println("Indique por favor bien el dijito")
		}
		fmt.Println("Ultima linea del bucle")
	}
	fmt.Println("Fuera del bucle")
}
