// Faça um Programa que pergunte quanto você ganha por hora e o número de horas trabalhadas no mês. Calcule e mostre o total do seu salário no referido mês.
package main

import "fmt"

func main() {

	var valorHora, horasTrabalhadas, totalSalario float64

	fmt.Println("Qual valor você ganhar por hora de trabalho?")
	fmt.Scanln(&valorHora)
	fmt.Println("Quantas horas você trabalhou no mês? ")
	fmt.Scanln(&horasTrabalhadas)

	totalSalario = (valorHora * horasTrabalhadas)
	fmt.Printf("O valor do seu salário total é: %2.f", totalSalario)
}
