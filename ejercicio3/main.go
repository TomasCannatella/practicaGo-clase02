/*
Ejercicio 3 - A qué mes corresponde

Realizar una aplicación que reciba  una variable con el número del mes.

Según el número,

	imprimir el mes que corresponda en texto.

¿Se te ocurre que se puede resolver de distintas maneras?

  - Si, hay muchas formas de resolver este ejercicio, se me ocurren 3 maneras diferente de hacerlo

  - La primera, es utilizando un if-else, donde en una variable almaceno el numero del mes y despues
    con un if voy validando cada caso, en caso del que mes no sea valido entraria en la ultima condición.

  - La segunda es muy similar a la anterior solo con la diferencia que en vez de utilizar un if-else anidado
    utilizo un switch para ver a que mes corresponde el número ingresado y poder mostrarlo por pantalla.

  - La tercer opción es utilizar un map, donde utilizo como clave el nro del mes y como valor el nombre del mes
    que corresponde. Y despues lo que hago es validar que el número ingresado sea un número valido del mes y luego
    con ese número busco el mes dentro del map para mostrarlo por pantalla.

¿Cuál elegirías y por qué?
En mi caso elgiria la opcion del map. Ya que en este ejercicio utilizando las 3 formas que se mencionan arriba a nivel
de perfomance las tres serian muy similares. Pero elijo la tercer opcion ya que me parece la opcion mas prolija a
nivel de código, en el caso de que hicieramos algo similar donde tuvieramos que agregar mas valores seria mucho mas
facil de hacerlo ya que lo tendriamos que agregar dentro del map sin modificar la lógica del codigo y seria mucho
mas facil de escalar.

Ej: 7, Julio.

Nota: Validar que el número del mes, sea correcto.
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	var monthNumber uint
	var continueProgram string

	/* Solucion con if else */
	// for {
	// 	fmt.Println("Ingrese un número del mes")
	// 	fmt.Scanf("%d", &monthNumber)

	// 	if monthNumber == 1 {
	// 		fmt.Println("El mes que corresponde a este número es: Enero")
	// 	} else if monthNumber == 2 {
	// 		fmt.Println("El mes que corresponde a este número es: Febrero")
	// 	} else if monthNumber == 3 {
	// 		fmt.Println("El mes que corresponde a este número es: Marzo")
	// 	} else if monthNumber == 4 {
	// 		fmt.Println("El mes que corresponde a este número es: Abril")
	// 	} else if monthNumber == 5 {
	// 		fmt.Println("El mes que corresponde a este número es: Mayo")
	// 	} else if monthNumber == 6 {
	// 		fmt.Println("El mes que corresponde a este número es: Junio")
	// 	} else if monthNumber == 7 {
	// 		fmt.Println("El mes que corresponde a este número es: Julio")
	// 	} else if monthNumber == 8 {
	// 		fmt.Println("El mes que corresponde a este número es: Agosto")
	// 	} else if monthNumber == 9 {
	// 		fmt.Println("El mes que corresponde a este número es: Septiembre")
	// 	} else if monthNumber == 10 {
	// 		fmt.Println("El mes que corresponde a este número es: Octubre")
	// 	} else if monthNumber == 11 {
	// 		fmt.Println("El mes que corresponde a este número es: Noviembre")
	// 	} else if monthNumber == 12 {
	// 		fmt.Println("El mes que corresponde a este número es: Diciembre")
	// 	} else {
	// 		fmt.Println("El número de mes ingresado no existe")
	// 	}

	// 	fmt.Println("Desea Ingresar otro número de mes (S/N)")
	// 	fmt.Scanf("%s", &continueProgram)

	// 	if strings.ToUpper(continueProgram) == "N" {
	// 		break
	// 	}
	// }

	/* Solución con switch */
	// for {
	// 	fmt.Println("Ingrese un número del mes")
	// 	fmt.Scanf("%d", &monthNumber)

	// 	switch monthNumber {
	// 	case 1:
	// 		fmt.Println("El mes que corresponde a este número es: Enero")
	// 	case 2:
	// 		fmt.Println("El mes que corresponde a este número es: Febrero")
	// 	case 3:
	// 		fmt.Println("El mes que corresponde a este número es: Marzo")
	// 	case 4:
	// 		fmt.Println("El mes que corresponde a este número es: Abril")
	// 	case 5:
	// 		fmt.Println("El mes que corresponde a este número es: Mayo")
	// 	case 6:
	// 		fmt.Println("El mes que corresponde a este número es: Junio")
	// 	case 7:
	// 		fmt.Println("El mes que corresponde a este número es: Julio")
	// 	case 8:
	// 		fmt.Println("El mes que corresponde a este número es: Agosto")
	// 	case 9:
	// 		fmt.Println("El mes que corresponde a este número es: Septiembre")
	// 	case 10:
	// 		fmt.Println("El mes que corresponde a este número es: Octubre")
	// 	case 11:
	// 		fmt.Println("El mes que corresponde a este número es: Noviembre")
	// 	case 12:
	// 		fmt.Println("El mes que corresponde a este número es: Diciembre")
	// 	default:
	// 		fmt.Println("El número de mes ingresado no existe")
	// 	}

	// 	fmt.Println("Desea Ingresar otro número de mes (S/N)")
	// 	fmt.Scanf("%s", &continueProgram)

	// 	if strings.ToUpper(continueProgram) == "N" {
	// 		break
	// 	}
	// }

	/* Solucion con map */
	var months = map[int]string{1: "Enero",
		2:  "Febrero",
		3:  "Marzo",
		4:  "Abril",
		5:  "Mayo",
		6:  "Junio",
		7:  "Julio",
		8:  "Agosto",
		9:  "Septiempre",
		10: "Octubre",
		11: "Noviembre",
		12: "Diciembre"}

	for {
		fmt.Println("Ingrese un número del mes")
		fmt.Scanf("%d", &monthNumber)

		if selectedMonth, ok := months[int(monthNumber)]; !ok {
			fmt.Println("El número de mes ingresado no existe")
		} else {
			fmt.Println("El mes que corresponde a este número es: ", selectedMonth)
		}

		fmt.Println("Desea Ingresar otro número de mes (S/N)")
		fmt.Scanf("%s", &continueProgram)

		if !(strings.ToUpper(continueProgram) == "S") {
			break
		}
	}

	fmt.Println("La ejecución del programa finalizo...")
}
