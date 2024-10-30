package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", true, true)           //Boolean
	fmt.Printf("Type: %T Value: %v\n", "Adriano", "Adriano") //String
	fmt.Printf("Type: %T Value: %v\n", 10, 10)               //Int
	fmt.Printf("Type: %T Value: %v\n", "10", "10")           //String
	fmt.Printf("Type: %T Value: %v\n", 1.474, 1.474)         //Float
}
