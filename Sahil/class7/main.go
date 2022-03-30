package main

import "fmt"

type apple func(int, int, string) (int, error)

func text(a apple) {
	a(5, 5, "")
}

func main() {

	var x apple = func(i1, i2 int, s string) (int, error) {
		fmt.Println(i1, i2, s)
		return 0, nil
	}
	text(x)

	var s = func() { //first class function
		fmt.Println("Hey hello")
	}

	addInts(1, 2, 3, 4, 5, 6)
	addInts(1)
	addInts(1, 3, 34, 4, 5, 6, 4, 7, 8, 8, 9, 9, 7, 6, 5)

	a := []int{1, 2, 34, 54, 6, 67, 5, 4, 43, 5, 6, 7, 4, 4343, 5, 6, 76, 5, 4, 433, 3}

	addInts(a...) //unpack
	s()
	var b int = 10
	dummy(b)
	fmt.Println(b)
	s()
	fmt.Println("Pass by reference")

	dummyRef(&b)
	fmt.Println(b)

	//anonymmous function
	func() {
		a[4] = 100000
		fmt.Println("Anonymous ", a)

	}()

	//First class functions
	s()

	var ps = func(i int) {
		fmt.Println(a, i)
	}

	printer(ps)

	ret := logic()
	ret()

}

func printer(fn func(int)) {
	fn(100000000)
}

func logic() func() {
	return func() {
		fmt.Println("logic function")
	}
}

func dummyRef(i *int) {
	*i += 10
	fmt.Println(*i)
}

func dummy(i int) {
	i += 10
	fmt.Println(i)
}

func addInts(a ...int) int { //pack   //pass by value
	var sum int
	for _, val := range a {
		sum += val
	}
	return sum
}
