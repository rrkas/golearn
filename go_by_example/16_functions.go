package main

import (
	"fmt"
	"math"
)

func main() {
	foo()
	circle_area(2.56)
	calculate(2, 6.7, 5)
	val := advanced_calculate(2, 4.8, 7)
	fmt.Printf("advanced_calculate() = %.2f \n", val)
	val1, val2, val3 := compute(6, 2.7, 1.4, "energy")
	fmt.Printf("val1=%.2f, val2=%.2f, val3=\"%s\" \n", val1, val2, val3)
	result := add(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11)
	fmt.Printf("add() = %d \n", result)
	closure_func("testing")
	ret := fibonacci(15)
	fmt.Printf("fibonacci() = %d \n", ret)
}

// a simple function
func foo() {
	fmt.Println("foo() was called")
}

// a function with a parameter
func circle_area(r float64) {
	area := math.Pi * math.Pow(r, 2)
	fmt.Printf("Circle area (r=%.2f) = %.2f \n", r, area)
}

// a function with parameters
func calculate(a, b, c float64) {
	result := a * b * c
	fmt.Printf("a=%.2f, b=%.2f, c=%.2f = %.2f \n", a, b, c, result)
}

// a function with parameters and a return value
func advanced_calculate(a, b, c float64) float64 {
	result := a * b * c
	return result
}

// a function with parameters and multiple return values
func compute(a, b, c float64, name string) (float64, float64, string) {
	result1 := a * b * c
	result2 := a + b + c
	result3 := result1 + result2
	newName := fmt.Sprintf("%s value = %.2f", name, result3)
	return result1, result2, newName
}

// a function with zero or more parameters and a return value
func add(numbers ...int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}
	return result
}

// a closure function
func closure_func(name string) {
	hoo := func(a, b int) {
		result := a * b
		fmt.Printf("hoo() = %d \n", result)
	}
	joo := func(a, b int) int {
		return a*b + a
	}
	fmt.Printf("closure_func(%s) was called\n", name)
	hoo(2, 3)
	val := joo(5, 8)
	fmt.Printf("val from joo() = %d \n", val)
}

// a recursion function
func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return (fibonacci(n-1) + fibonacci(n-2))
}
