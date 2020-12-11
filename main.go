package main

import "fmt"

// declare parameter and return data type
func multiply(a int, b int) int {
	return a * b
}

// if, parameter has same data type, u can set paramere data type as follow
func multiply2(a, b int) int {
	return a * b
}

func main() {
	fmt.Println(multiply(2, 2))
	fmt.Println(multiply2(4, 3))
}
