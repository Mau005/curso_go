package capituloUno

import "fmt"

func Operadores() {
	var texto string = "hola"
	var texto1 string = "hola"

	resultado := texto == texto1 // false
	//resultado = num1 < num2 // true
	/*
		== comparacion absoluta seria si este si o si es igual al otro
		!= si es distinto del valor de la izquierda al de la derecha
		< menor
		> mayor
		<= menor o igual
		>= mayor o igual

		Operadores Logicos ---
	*/
	fmt.Println(resultado)
}
