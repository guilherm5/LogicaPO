package main

import "fmt"

func main() {
	var base float64
	var altura float64

	fmt.Println("Calculo Area do retangulo")
	fmt.Println("Qual a base deste retangulo ?")

	fmt.Scanln(&base)

	fmt.Println("Qual a altura deste retangulo ?")

	fmt.Scanln(&altura)

	var calculoArea = base * altura

	fmt.Println("A area deste retangulo é:", calculoArea)

	fmt.Println("Agora vamos ao calculo do perimetro:")

	var calculaPerimetro = (2 * (base + altura))

	fmt.Println("O calculo do perimetro é:", calculaPerimetro)

}
