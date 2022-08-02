package main

import "fmt"

func main() {
	var edad int
	fmt.Scan(&edad)

	//&puntero *valor

	switch {
	case edad >= 18 && edad <= 65:
		fmt.Println("Puede entrar")
	case edad < 18:
		fmt.Println("Eres menor")
	case edad > 65:
		fmt.Println("eres muy mayor")
	}
}
