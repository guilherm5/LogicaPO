package main

import "fmt"

func main() {
	var idadeEmDias int

	fmt.Println("Bem vindo a sua calculadora de idade")
	fmt.Println("Para começar, digite sua idade em dias")
	fmt.Scanln(&idadeEmDias)

	var idadeEmAnos = idadeEmDias / 365
	var meses = (idadeEmDias % 365) / 30
	var dias = (idadeEmDias % 365) % 30

	fmt.Printf("A idade é: %d anos, %d meses e %d dias\n", idadeEmAnos, meses, dias)
}
