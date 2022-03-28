package main

import "fmt"

func pointers() {
	var i int = 10
	var c *int = &i

	fmt.Println(i, c, *c)

	i = 20
	fmt.Println(i, c, *c)

	*c = *c + 10
	fmt.Println(i, c, *c)
}
