package main

import (
	"bufio"
	"fmt"
	"os"
)

func userio() {
	// var i string
	// fmt.Scanln(&i)
	// fmt.Println("userio ", i)

	// var d int
	// fmt.Scanf("%d", &d)
	// fmt.Println("userio ", d)

	bufscan := bufio.NewScanner(os.Stdin)
	// for bufscan.Scan() {
	// 	i = bufscan.Text()
	// 	if i == "exit" {
	// 		break
	// 	}
	// }
	bufscan.Scan()
	i := bufscan.Text()
	fmt.Println("userio ", i)
}
