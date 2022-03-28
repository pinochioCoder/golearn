package main

import "fmt"

func main() {
	a := 0
	a, b := addNum(4, 5)
	fmt.Println("", a, b)
	// fmt.Println("", addNum(5, 6))
	// fmt.Println("", addNum(1, 1))

}

func addNum(a, b int) (c, d int) {
	c = a + b
	return
}
