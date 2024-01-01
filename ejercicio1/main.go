/*
La Real Academia Española quiere saber cuántas letras tiene una palabra
y luego tener cada una de las letras por separado para deletrearla. Para eso tendrán que:

- Crear una aplicación que reciba una variable con la palabra e imprimir la cantidad de letras que contiene la misma.
- Luego, imprimir cada una de las letras.
*/
package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("Sistema de LA REAL ACADEMIA ESPAÑOLA")
	fmt.Println("- Imprime la cantidad de letras que tiene la palabra")
	fmt.Println("- Imprime la c/u de las letras")

	var word string
	var continueProgram string
	for {
		fmt.Println("Ingrese una palabra")

		fmt.Scanf("%s", &word)

		fmt.Println("La longitud de su palabra es de: ", len(word))

		fmt.Print("Las letras de su palabra son: ")
		for _, char := range word {
			fmt.Print(string(char), " ")
		}
		fmt.Println("")

		fmt.Println("Desea volver a ingresar otra palabra S/N")

		fmt.Scanf("%s", &continueProgram)
		if !(strings.ToUpper(continueProgram) == "S") {
			break
		}
	}
	fmt.Println("La ejecución del programa ha finalizado...")
}
