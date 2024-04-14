package main

import (
	"fmt"
)

func main() {

	var b int
	a := 1
	fmt.Printf("Enter a number")
	fmt.Scanf("%d", &b) //Read input number
	soma := a + b
	fmt.Println("My return:", soma) //Write result of sum

}
