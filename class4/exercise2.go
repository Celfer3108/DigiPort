package main
 type Funcionario struct {
nome    string
idade   int
função  string
salario int
 }

import (
	"fmt"
)

func main() {

	a := Funcionario{"Maria", 39, "Assistente", 3000.00}
	b := Funcionario{"Madonna", 48, "Professora", 5000.00}
	c := Funcionario{"Priscila", 25, "Engenheira", 15000.00}

	var funcionario [3]Funcionario

	funcionario[0] = a
	funcionario[1] = b
	funcionario[2] = c

	fmt.Println(funcionario[1])
}
