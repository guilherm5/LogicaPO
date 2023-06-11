package main

import "fmt"

func main() {
	var temp float64

	fmt.Println("Bem vindo a sua calculadora de Fahrenheit x Celsius")
	fmt.Println("Digite uma temperatura em Fahrenheit:")
	fmt.Scanln(&temp)

	var converte = (temp - 32) / 1.8

	fmt.Println("A temperatura em Fahrenheit é:", converte, "C")
}
