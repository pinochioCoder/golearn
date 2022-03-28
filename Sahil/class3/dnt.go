package main

import (
	"fmt"
	"time"
)

func dnt() {
	fmt.Println(time.Now())
	fmt.Println(time.Now().Add(-time.Hour * 1))
	fmt.Println(time.Now().Format("02-01-2006 03:04 PM"))
}
