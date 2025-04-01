package main

import (
	"fmt"
)

func main() {
	// Ejercicio 1: Calcular el perímetro de un triángulo
	fmt.Println("Ejercicio 1")
	perimetro := PerimetroTriangulo(10, 20, 30)
	fmt.Println(perimetro)

	// Ejercicio 2: Calcular el monto total a pagar a un empleado
	fmt.Println("Ejercicio 2")

	var nombre string
	var horasTrabajadas int
	var valorHora float64

	fmt.Print("Ingrese el nombre del empleado: ")
	fmt.Scanln(&nombre)
	fmt.Print("Ingrese las horas trabajadas: ")
	fmt.Scanln(&horasTrabajadas)
	fmt.Print("Ingrese el valor por hora: ")
	fmt.Scanln(&valorHora)

	empleado := Empleado{
		Nombre:          nombre,
		HorasTrabajadas: horasTrabajadas,
		ValorHora:       valorHora,
	}

	montoTotal := empleado.CalcularMontoTotal()
	fmt.Printf("El monto total a pagar a %s es: %.2f\n", empleado.Nombre, montoTotal)
}
