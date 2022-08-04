package main

import "fmt"

func main() {
	edad := 18
	sexo := "Mujer"
	//Club que solo las personas entre 18 y 64 años pueden entrar
	if edad >= 18 && edad <= 64 {
		fmt.Println("Eres mayor puedes entrar")
		if sexo == "Mujer" {
			fmt.Println("Tienes un descuento del 65%")
		}
	} else {
		fmt.Println("No Cumples los requisitos minimo para poder entrar")
	}

}
