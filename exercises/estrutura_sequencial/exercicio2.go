// Faça um Programa que peça um número e então mostre a mensagem O número informado foi [número]
package main

import "fmt"

func main() {
	var num int

	fmt.Println("Digite um número: ")
	fmt.Scanln(&num)
	fmt.Printf("O número informado foi %v", num)
}
