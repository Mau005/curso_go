package main

import "fmt"

func main() {
	var texto string = "Hola Mundo"

	for i := 0; i < 10; i++ {
		fmt.Printf("%c\n", texto[i])
	}
}
