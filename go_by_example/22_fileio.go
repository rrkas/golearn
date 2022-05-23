package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// write data into a file
	fmt.Println("writing data into a file")
	writeFile("welcome to go")
	readFile()

	fmt.Println()

	// read data from a file
	fmt.Println("reading data from a file")
	readFile()
}

func writeFile(message string) {
	bytes := []byte(message)
	ioutil.WriteFile("../temp/testgo.txt", bytes, 0644)
	fmt.Println("created a file")
}

func readFile() {
	data, _ := ioutil.ReadFile("../temp/testgo.txt")
	fmt.Println("file content:")
	fmt.Println(string(data))
}
