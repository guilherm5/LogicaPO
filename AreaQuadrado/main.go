package main

import "fmt"

func main() {
	var ladoQuadrado float64

	fmt.Println("Calculo Area do quadrado")
	fmt.Println("Qual a medida de um dos lados do quadrado ?")

	fmt.Scanln(&ladoQuadrado)

	var calculoArea = ladoQuadrado * ladoQuadrado

	fmt.Println("A area deste quadrado é:", calculoArea)

	fmt.Println("Agora vamos ao calculo do perimetro:")

	var calculaPerimetro = 4 * ladoQuadrado

	fmt.Println("O calculo do perimetro é:", calculaPerimetro)
}
