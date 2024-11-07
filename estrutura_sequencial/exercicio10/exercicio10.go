// 10 - Faça um Programa que peça a temperatura em graus Celsius, transforme e mostre em graus Fahrenheit.
package main

import (
	"fmt"
)

func main() {

	var temperaturaC, temperaturaF float64

	fmt.Println("Qual a temperatura em celsius? ")
	fmt.Scanln(&temperaturaC)

	temperaturaF = (temperaturaC * 1.8) + 32
	fmt.Printf("A temperatura em celsius: %.2f convertida para Farenheit fica: %.2f ", temperaturaC, temperaturaF)
}
