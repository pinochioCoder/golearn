package main

import "fmt"

func tmpfunc() {
	defer fmt.Println("I am first")
	defer fmt.Println("I am second")
	defer fmt.Println("I am third")
	defer fmt.Println("I am fourth")

	var a int = 10
	var b int = 20
	c := a + b
	fmt.Println("Sum is ", c)

}
