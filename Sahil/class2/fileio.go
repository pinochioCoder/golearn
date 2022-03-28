package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func fileio() {

	filedata, _ := ioutil.ReadFile("text.txt")
	fmt.Println(string(filedata))

	data := []byte("this is dummy file data that we will write to the file")
	ioutil.WriteFile("dummyFile.txt", data, os.ModeAppend)

}
