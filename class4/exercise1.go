package main

import (
	"fmt"
)

func main() {

	func RemoveVowels(word string) string {
	for _, c := range [string] {"a","e","i","o","u","A","E","I","O","U"} {
	ord = strings.Replace (word, c, "", -1)
		}
			return word
	}
	func main() {
	var word = ""
	fmt.Println ("Enter a word: ")
	fmt.Scanf("%s", &word)  //read input from STDIN
	fmt.Println("My entry: ", word) //write output to STDOUT
	newWord := RemoveVowels(word)
	fmt.Println("My Return", newWord) //write output to STDOUT
	}
}