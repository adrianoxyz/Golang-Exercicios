// Faça um Programa que peça o raio de um círculo, calcule e mostre sua área.
package main

import (
	"fmt"
)

func main() {
	var r, area, raioAoQuadrado float64
	const Pi = 3.14

	fmt.Println("Digite o raio a ser calculado: ")
	fmt.Scanln(&r)
	raioAoQuadrado = r * r
	area = Pi * raioAoQuadrado

	fmt.Printf("O valor da área é: %.2f", area)

}
