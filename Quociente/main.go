package main

import "fmt"

func main() {
	var numero1 float64
	var numero2 float64

	fmt.Println("Bem vindo a sua calculadora de divisão")
	fmt.Println("Para começar digite um numero")
	fmt.Scanln(&numero1)
	fmt.Println("Digite outro numero")
	fmt.Scanln(&numero2)

	var resultadoDivisao = numero1 / numero2
	var restoDivisao = int(numero1) % int(numero2)

	fmt.Println("O resultado da divisão é:", resultadoDivisao)
	fmt.Println("O resultado do resto da divisão é:", restoDivisao)
}
