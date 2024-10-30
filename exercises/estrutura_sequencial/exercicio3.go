// Faça um Programa que peça dois números e imprima a soma.

package main

import "fmt"

func main() {

	var num1, num2, soma int

	fmt.Println("Digite o 1° número: ")
	fmt.Scanln(&num1)
	fmt.Println("Digite o 2° número: ")
	fmt.Scanln(&num2)
	soma = num1 + num2
	fmt.Println("A soma de", num1, "e", num2, "foi: ", soma)

	// Tipos:
	// Boolean (True/False)
	// String - Sequência de Bytes
	// int
	// Float (float64/float32) - decimal

}
