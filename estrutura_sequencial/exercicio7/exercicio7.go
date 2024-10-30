// Faça um Programa que calcule a área de um quadrado, em seguida mostre o dobro desta área para o usuário.
package main

import "fmt"

func main() {
	var r, area, raioAoQuadrado, areaAoQuadrado float64
	const Pi = 3.14

	fmt.Println("Digite o raio a ser calculado: ")
	fmt.Scanln(&r)
	raioAoQuadrado = r * r
	area = Pi * raioAoQuadrado
	areaAoQuadrado = area * area

	fmt.Printf("O valor da área é: %.2f", areaAoQuadrado)
}
