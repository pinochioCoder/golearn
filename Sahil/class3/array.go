package main

import (
	"fmt"
	"reflect"
	"sort"
)

func array() {
	//[4,5,6,7,8] -> data
	//[0,1,2,3,4] -> data location
	i := []int{4, 5, 6, 7, 8}
	fmt.Println(len(i))
	fmt.Println(i[len(i)-1])
	//len ->
	//index ->

	for idx, val := range i {
		fmt.Printf("index: %d, Value: %d \n", idx, val)
	}
	fmt.Println(i)

	var slc []int
	fmt.Printf("%T \n", slc)
	fmt.Println(reflect.TypeOf(slc).String())

	slc = i[:2]

	fmt.Println("Values ", slc)
	fmt.Println("length: ", len(slc))
	fmt.Println("capacity: ", cap(slc))

	slc[1] = 100

	fmt.Println(i)
	fmt.Println(slc)

	i[0] = 200

	fmt.Println(i)
	fmt.Println(slc)

	sort.Ints(i)
	fmt.Println(i)

}
