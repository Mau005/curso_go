package main

import "fmt"

func main() {
	/*
		+ sumar
		- restar
		* multiplicar
		/ dividir
		% resto de
		() parentesis
		primero sacara el calculo de ()
		luego la multiplicacion mas cercana a la izquieda
		y luegos las sumas o restas
	*/

	var num1 int = 10
	var num2 int = 10

	resultado := num1 % num2
	fmt.Println(resultado)

	fmt.Println("--------------")

	resultado2 := 8 + 3*(1+2)%5
	/*
		8 + 3 * 3 %4
		8 + 9 %5
		12
	*/

	fmt.Println(resultado2)

}
