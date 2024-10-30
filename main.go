package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, Go Essentials!")   //String
	fmt.Printf("Saída formatada: %v\n", 2) // Saída formatada com printf e utilizando String e números.

	fmt.Println(2) //Inteiros

	fmt.Println(2.2) //Float

	fmt.Println(2 + 5) //Expressões

	// Booleanos
	fmt.Println(true)
	fmt.Println(false)
	fmt.Println(strings.Split("Adriano", ""))
}

// Pasta == folders podem ter múltiplos arquivos.
// Pasta pode ter múltiplas funções dentro dela (espalhada em múltiplos arquivos)
// Toda pasta é um pacote
// Funções pertencem a um pacote
