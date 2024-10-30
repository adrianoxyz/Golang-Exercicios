// Faça um Programa que peça as 4 notas bimestrais e mostre a média.
package main

import (
	"fmt"
)

func main() {

	var nota1, nota2, nota3, nota4, media float64

	fmt.Println("Digite a nota do 1° bimestre: ")
	fmt.Scanln(&nota1)
	fmt.Println("Digite a nota do 2° bimestre: ")
	fmt.Scanln(&nota2)
	fmt.Println("Digite a nota do 3° bimestre: ")
	fmt.Scanln(&nota3)
	fmt.Println("Digite a nota do 4° bimestre: ")
	fmt.Scanln(&nota4)

	media = (nota1 + nota2 + nota3 + nota4) / 4
	fmt.Println("A média do aluno foi: ", media)
}
