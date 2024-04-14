package main

import (
	"fmt"
)

func main() {
	fmt.Println("Qual seu nome?")
	var nome string
	fmt.Scanf("%s", &nome)
	fmt.Println("Bem-vindo, ", nome)
	fmt.Printf("Bem vindo, %s", nome)
}
