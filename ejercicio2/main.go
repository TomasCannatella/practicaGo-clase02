/*
Ejercicio 2 - Préstamo

Un banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los mismos.
Para ello tiene ciertas reglas para saber a qué cliente se le puede otorgar.
Solo le otorga préstamos a clientes cuya edad sea mayor a 22 años,
se encuentren empleados
y tengan más de un año de antigüedad en su trabajo.

Dentro de los préstamos que otorga
no les cobrará interés a los que posean un sueldo superior a $100.000.

Es necesario realizar una aplicación que reciba
estas variables y que imprima un mensaje de acuerdo a cada caso.

Tip: tu código tiene que poder imprimir al menos 3 mensajes diferentes.
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		age             uint
		isEmployee      string // puede tomar Si o No (S/N)
		antiquity       uint
		salary          uint
		interest        uint
		loan            uint
		continueProgram string
	)

	for {
		fmt.Println("Ingrese la edad del cliente")
		fmt.Scanf("%d", &age)

		if !(age > 22) {
			fmt.Println("El cliente debe ser mayor a 22 años para pedir el prestamo")
			return
		}

		fmt.Println("Indique si es empleado S/N")
		fmt.Scanf("%s", &isEmployee)

		if !(strings.ToUpper(isEmployee) == "S") {
			fmt.Println("El cliente debe ser empleado para pedir el prestamo")
			return
		}

		fmt.Println("Ingrese la antiguedad del empleado")
		fmt.Scanf("%d", &antiquity)

		if !(antiquity > 1) {
			fmt.Println("El cliente debe tener mas de un año de antiguedad para pedir el prestamo")
			return
		}

		fmt.Println("Ingrese el sueldo del cliente")
		fmt.Scanf("%d", &salary)

		fmt.Println("Ingrese el monto del prestamo")
		fmt.Scanf("%d", &loan)

		fmt.Println("Ingrese el interes del prestamo")
		fmt.Scanf("%d", &interest)

		if !(salary > 100000) {
			loan += (loan * interest) / 100
			fmt.Println("El banco le otorga un prestamo de $", loan, " con un interes del ", interest, "%")
			return
		}

		fmt.Println("El banco le otorgo un prestamo de: $", loan)

		fmt.Println("Desea volver a pedir otro prestamo")
		fmt.Scanf("%s", &continueProgram)
		if !(strings.ToUpper(continueProgram) == "S") {
			break
		}
	}
}
