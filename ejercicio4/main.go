/*
Ejercicio 4 - Qué edad tiene...

Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados.
Según el siguiente mapa, ayudá a imprimir la edad de Benjamin.
var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
Por otro lado también es necesario:
- Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
- Eliminar a Pedro del mapa.
- Saber cuántos de sus empleados son mayores de 21 años.
*/
package main

import "fmt"

func main() {

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	fmt.Println("La edad de Benjamin es: ", employees["Benjamin"])

	fmt.Println("Agregando al empleado Federico....")
	employees["Federico"] = 25
	fmt.Println("Eliminando a Pedro del mapa...")
	delete(employees, "Pedro")

	fmt.Println("Los empleados que son mayores a 21 son:")
	for name, age := range employees {
		if age > 21 {
			fmt.Println("- ", name, " => ", age, " años")
		}
	}

}
