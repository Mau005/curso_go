package main

import "fmt"

func main() {
	a := 0
	b := 0

	for i := 0; i < 10; i++ {
		if true {
			a := 1
			b = 1

			a++
			b++
		}
	}

	fmt.Printf("A: %v B: %v", a, b)

}
