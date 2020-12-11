package main

import (
	"fmt"
	"strings"
)

// declare parameter and return data type
func multiply(a int, b int) int {
	return a * b
}

// if, parameter has same data type, u can set paramere data type as follow
func multiply2(a, b int) int {
	return a * b
}

// return multiple value
func lenAndUpperOfString(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

// get all parameters
func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	fmt.Println(multiply(2, 2))
	fmt.Println(multiply2(3, 4))
	name := "JohnMark"
	fmt.Println(lenAndUpperOfString(name))

	nameLen, upperName := lenAndUpperOfString(name)

	fmt.Printf("My Name is %s , Name Length is %d\n", upperName, nameLen)

	repeatMe("a", "b", "c", "d", "Zip")
}
