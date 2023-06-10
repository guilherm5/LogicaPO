package main

import "fmt"

func main() {
	var lado1 float64
	var lado2 float64
	var lado3 float64

	fmt.Println("Calculo perimetro do triangulo")

	fmt.Println("Qual a medida do primeiro lado do triangulo ?")
	fmt.Scanln(&lado1)

	fmt.Println("Qual a medida do segundo lado do triangulo ?")
	fmt.Scanln(&lado2)

	fmt.Println("Qual a medida do terceiro lado do triangulo ?")
	fmt.Scanln(&lado3)

	var calculaPerimetro = lado1 + lado2 + lado3

	fmt.Println("O calculo do perimetro é:", calculaPerimetro)
}
