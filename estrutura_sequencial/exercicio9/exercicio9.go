// Faça um Programa que pergunte quanto você ganha por hora e o número de horas trabalhadas no mês. Calcule e mostre o total do seu salário no referido mês.
package main

import (
	"fmt"
)

func main() {

	var valorHora, horasTrabalhadasMes, valorSalario float64

	fmt.Println("Qual valor vocÊ recebe por hora?")
	fmt.Scanln(&valorHora)
	fmt.Println("Quantas horas você trabalha no mês?")
	fmt.Scanln(&horasTrabalhadasMes)

	valorSalario = valorHora * horasTrabalhadasMes
	fmt.Printf("O valor do seu salário é: %2.f", valorSalario)
}
