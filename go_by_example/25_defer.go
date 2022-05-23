package main

import (
	"fmt"
)

func main() {
	demoPanic()
}
func demoPanic() {
	fmt.Println("----demo error handling---")
	defer func() {
		fmt.Println("do something")
	}()
	panic("this is a panic from demoPanic()")
	fmt.Println("This message will never show")
}
