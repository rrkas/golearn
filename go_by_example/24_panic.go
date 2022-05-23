package main

import (
	"fmt"
)

func main() {
	demoPanic()
}
func demoPanic() {
	fmt.Println("before panic")

	panic("this is a panic from demoPanic()")

	fmt.Println("This message will never show")
}
