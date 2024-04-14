package main

import (
	"fmt"
)

func main() {

	//decreasing
	for contador := 10; contador > 0; contador-- {
		fmt.Println("Count down", contador)
	}
	fmt.Println("Happy New Year")
	//increasing
	for contador := 0; contador > 10; contador++ {
		fmt.Println("Count up", contador)
	}
	fmt.Println("Arrasou")
}
