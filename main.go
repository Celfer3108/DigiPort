package main

import "fmt"

func main() {
	fmt.Println("Hello world")
}

// exercicio 1
// fmt.Println("Qual seu nome?")
// var nome string
// fmt.Scanf("%s", &nome)
// fmt.Println("Bem-vindo, ", nome)
// fmt.Printf("Bem vindo, %s", nome)

// exercicio 2
// var b int
// a := 1
// fmt.Printf("Enter a number")
// fmt.Scanf ("%d", &b) //Read input number
// soma := a + b
// fmt.Println ("My return:", soma) //Write result of sum

// exercício 3
//     var num int
//     fmt.Println ("Qual sua idade?")
//     fmt.Scanf ("%d", &num)

//     if num > 0 {
// 	   fmt.Println ("Positivo!")
// 	} else if num < 0 {
// 	   fmt.Println ("Negative!")
// 	} else {
// 	   fmt.Println("Zero")
// 	}

// exercício 4
// //decreasing
// for contador :=10; contador > 0; contador-- {
// 	fmt.Println ("Count down", contador)
// }
// fmt.Println("Happy New Year")
// //increasing
// for contador :=0; contador > 10; contador ++ {
// 	fmt.Println("Count up", contador)
// }
// fmt.Println("Arrasou")

//exercício 5
// 	//import package
// func RemoveVowels(word string) string {
// 	for _, c := range [string] {"a","e","i","o","u","A","E","I","O","U"} {
// 		word = strings.Replace (word, c, "", -1)
// 	}
// 		return word
// }
// func main() {
// 	var word = ""
// 	fmt.Println ("Enter a word: ")
// 	fmt.Scanf("%s", &word)  //read input from STDIN
// 	fmt.Println("My entry: ", word) //write output to STDOUT
// 	newWord := RemoveVowels(word)
// 	fmt.Println("My Return", newWord) //write output to STDOUT
//  }
