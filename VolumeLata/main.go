package main

import "fmt"

func main() {
	var raioCilindro float64
	var alturaCilindro float64
	fmt.Println("Bem vindo a sua calculadora de volume de cilindro")

	fmt.Println("Forneça o raio do cilindro")
	fmt.Scanln(&raioCilindro)
	fmt.Println("Forneça a altura do cilindro")
	fmt.Scanln(&alturaCilindro)

	var raioCalculado = raioCilindro * raioCilindro
	var calculoVolume = 3.1 * raioCalculado * alturaCilindro

	fmt.Println("O volume do seu Cilindro é:", calculoVolume)
}
