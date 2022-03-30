package main

import (
	"bufio"
	"fmt"
	"os"
)

func bufscanEx() {
	bufscn := bufio.NewScanner(os.Stdin)

	for bufscn.Scan() {
		txt := bufscn.Text()
		if txt == "exit" {
			fmt.Println("BYE NOW")
			break
		}
	}
}
