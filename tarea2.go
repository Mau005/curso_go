package main

import "fmt"

func main() {
	/*
		crea una aplicacion que pueda iniciar seccion
		dentro de el programa pueda crear un comentario
		y pueda salir del programa

	*/

	var condicion bool = true
	var usuario string = "mau"
	var contraseña string = "1234"
	var comentario string

	for condicion {
		var opcion int
		fmt.Println("[1] Iniciar Cuenta")
		fmt.Println("[2] salir del programa")
		fmt.Println("Deme la alternativa: ")
		fmt.Scan(&opcion)

		switch opcion {
		case 1:
			usuarioTem := ""
			fmt.Println("Indicame el usuario: ")
			fmt.Scan(&usuarioTem)
			contraseñaTemp := ""
			fmt.Println("Indicame la contraseña: ")
			fmt.Scan(&contraseñaTemp)
			if usuarioTem == usuario && contraseñaTemp == contraseña {
				condicionNueva := true // generamos una condicion
				for condicionNueva {
					tempOpcion := 0 //Generamos una variable temporal
					fmt.Println("Aqui inicia nuestra app")
					fmt.Println("[1] Crear Comentario")
					fmt.Println("[2] Mostrar el Comentario")
					fmt.Println("[3] desconectarse")

					fmt.Scan(&tempOpcion)

					switch tempOpcion {
					case 1:
						fmt.Println("Indicame el comentario nuevo: ")
						fmt.Scan(&comentario)
						fmt.Println("El comentario se ha guardado con exito")
					case 2:
						fmt.Printf("\nEl comentario es: %v", comentario)
					case 3:
						fmt.Println("Se procede a descoenctarse el usuario")
						condicionNueva = false
						condicion = false
					default:
						fmt.Println("Ninguna alternativa correcta intentalo nuevamente")
					}
				}

			} else {
				fmt.Println("Contraseñas incorrecta, intentelo nuevamente")
			}

		case 2:
			fmt.Println("Muchas gracias por usar esta app.")
			condicion = false
		default:
			fmt.Println("Mal colocado la informacion")
		}
	}

}
