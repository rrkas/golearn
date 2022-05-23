package main

import "fmt"

func main() {
	var i int

	for i = 0; i < 5; i++ {
		fmt.Println(i)
	}

	for j := 5; j < 11; j++ {
		fmt.Println(j)
	}

	i = 0
	for i < 5 {
		fmt.Println(i)
		i++
	}
}
