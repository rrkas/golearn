package main

import(
	"fmt"
	"github.com/rrkas/golearn/go_by_example/simplemaths"
)

func main() {
	fmt.Println("access mymodule…")
	var c int
	c = simplemaths.Add(10, 6)
	fmt.Printf("add()=%d\n", c)
	c = simplemaths.Subtract(5, 8)
	fmt.Printf("subtract()=%d\n", c)
	c = simplemaths.Multiply(5, 3)
	fmt.Printf("multiply()=%d\n", c)
}
