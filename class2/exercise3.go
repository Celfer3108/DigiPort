package main

import (
	"fmt"
)

func main() {

	var num int
	fmt.Println("Qual sua idade?")
	fmt.Scanf("%d", &num)

	if num > 0 {
		fmt.Println("Positivo!")
	} else if num < 0 {
		fmt.Println("Negative!")
	} else {
		fmt.Println("Zero")
	}
}
