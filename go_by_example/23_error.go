package main

import (
	"fmt"
)

func main() {
	calculate()
}

func calculate() {
	fmt.Println("----demo error handling---")
	a := 10
	b := 0
	c := 0
	c = a / b
	fmt.Printf("result = %.2f \n", c)
	fmt.Println("Done")
}
