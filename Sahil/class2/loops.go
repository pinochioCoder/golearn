package main

import "fmt"

func loops() {

	for i := 0; i <= 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Print(" ", i)
	}

	for {
		fmt.Println(1)
		fmt.Println(1)
		break
	}

	var i int = 1
	for i < 5 {
		fmt.Print(" ", i)
		i++
	}

}
