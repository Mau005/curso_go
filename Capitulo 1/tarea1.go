package main

import "fmt"

func main() {
	/*
	   crear un programa con 3 variables edad, nombre, sexo, OK
	   mensajes filtrados entre 18 y 64años que opueda entrar OK
	   si tienes mas de 65 diga que no puede entrar
	   si tienes menos de 18 diga que es un menor para estas cosas OK
	   si es mujer tiene descuento del 65% OK
	   si es pedro y es hombre entre totalmente gratis OK
	*/
	var edad uint = 141
	var nombre string = "Pedro"
	var sexo string = "Mujer"

	if edad >= 18 && edad <= 65 {
		fmt.Println("Ingresando")

		if sexo == "Mujer" {
			fmt.Println("Tienes descuento del 65%")
		} else {
			if nombre == "Pedro" {
				fmt.Println("Usted entra gratis")
			} else {
				fmt.Println("Usted ingresa normal y tiene que pagar")
			}
		}
	} else if edad < 18 {
		fmt.Println("Usted es muy joven para el ingreso")
	} else if edad > 65 {
		fmt.Println("Usted esta muy viejo para estas cosas vallase a un asilo")
	}

	/*
		bloque es comparar el ingreso

		3 posivilidades
		entre 18 y 65
		menor a 18
		mayor 65
	*/

}
