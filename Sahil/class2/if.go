package main

import "fmt"

func ifs() {

	// AND &&
	//OR ||
	//NOT !

	// var days string = "tuesday"
	// var todayDate int = 10

	// if days != "monday" || todayDate != 10 {
	// 	fmt.Println("Monday")
	// } else if days == "tuesday" {
	// 	fmt.Println("tuesday")
	// } else {
	// 	fmt.Println("tuesday")
	// }

	// switch days {
	// case "monday", "wednesday", "friday":
	// 	fmt.Println("Holiday")
	// 	fallthrough
	// case "tuesday", "thursday", "saturday":
	// 	fmt.Println("workday")
	// case "sunday":
	// 	fmt.Println("sunday")
	// default:
	// }

	if v := sum(); v == 10 {
		fmt.Println("Sum found", v)
	}

}

func sum() int {
	return 5 + 5
}
