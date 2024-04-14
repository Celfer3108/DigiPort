package main

import (
	"fmt"
)

contatinhos := map[string]int{"Gui": 9999999999, "Paulinho": 8888888888, "Henrique": 777777777}
	fmt.Printf("%d", contatinhos["Henrique"])

	contactlist := make(map[string]string)

	contactlist["Gui"] = "Jogar bola"
	contactlist["Paulinho"] = "Pescar"
	contactlist["Henrique"] = "Nadar"

	for name, hobby := range contactlist {
		fmt.Printf("Name: %s, Hobby: %s\n", name, hobby)
	
}
