package main

import "fmt"

func main() {
	var contador int = 1

	for {
		fmt.Println("indicame un numero")
		fmt.Scan(&contador)

		if contador == 100 {
			break
		}

	}

}
