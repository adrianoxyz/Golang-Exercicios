// Faça um Programa que converta metros para centímetros.
package main

import (
	"fmt"
)

func main() {

	var metros, cm float64

	fmt.Println("Digite a quantidade de metros a ser convertido? ")
	fmt.Scanln(&metros)

	cm = (metros * 100)

	fmt.Printf("O valor de %v, convertido para centímetros, fica em: %v", metros, cm)
}
