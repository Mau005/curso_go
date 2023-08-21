package capituloUno

import "fmt"

func LenAndCap(array []int) {
	fmt.Println("La Longitud es: ", len(array))
	fmt.Println("La Capacidad es: ", cap(array))
}

func AsignarElementos(array [6]int) {
	array[0] = 10
	array[len(array)-1] = 10 //El ultimo digito del index
}

func AsigneNumeros(array *[6]int, rangeItems ...int) {
	aux := 0
	for _, value := range rangeItems {
		array[aux] = value
		aux++
		if len(array) <= aux {
			aux = 0
		}
	}
}

func AsigneNumerosSlice(mySlice *[]int) {
	*mySlice = append(*mySlice, 10)
	*mySlice = append(*mySlice, 15)
}
func AsigneNumerosSliceTwo(mySlice []int) []int {
	mySlice = append(mySlice, 10)
	mySlice = append(mySlice, 15)
	return mySlice
}

func TestSliceAndArray() {
	//Longitud y Capacidad
	MyArray := [6]int{1, 2, 3, 4, 5, 6}
	//AsigneNumeros(&MyArray, 10, 20, 30, 40, 50, 60, 10, 20, 324, 234, 324, 234, 324)
	//fmt.Println(MyArray)

	var MySlice []int
	/*
		for i := 0; i < 5; i++ {
			MySlice = append(MySlice, i*50)
		}
	*/
	LenAndCap(MySlice)
	LenAndCap(MyArray[:])
	MyArrayTwo := AsigneNumerosSliceTwo(MyArray[:])
	LenAndCap(MyArrayTwo)
	fmt.Println("MyArray: ", MyArrayTwo)
	AsigneNumerosSlice(&MySlice)
	LenAndCap(MySlice)
	fmt.Println("MySlice: ", MySlice)
}
