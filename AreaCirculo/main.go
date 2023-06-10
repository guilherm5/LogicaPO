package main

import "fmt"

func main() {
	var Circulo float64
	var raio float64

	fmt.Println("Calculo Area do circulo")
	fmt.Println("Qual o raio do circulo ?")

	fmt.Scanln(&Circulo)

	var calculoArea = Circulo * Circulo * (3.14)

	fmt.Println("A area deste circulo é:", calculoArea)

	fmt.Println("Agora vamos ao calculo do perimetro, digite o raio do circulo:")

	fmt.Scan(&raio)

	var calculaPerimetro = (2 * 3.14) * (raio) // 2 * pi * r, R é o raio, forneceremos o raio então

	fmt.Println("O calculo do perimetro do circulo é:", calculaPerimetro)
}
