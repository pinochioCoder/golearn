package main

import "fmt"

func runes() {
	var r rune
	fmt.Printf("%c \n", r)

	r = 'b'
	fmt.Printf("%c \n", r)
	r = 126
	fmt.Printf("%c \n", r)

	r = '!'
	fmt.Println(r)
}
