package main

import (
	"fmt"

	"github.com/LoveAndCode/golearn/echo"
	"github.com/LoveAndCode/golearn/functions"
)

func superAdd(numbers ...int) int {
	var sum int
	for index, number := range numbers {
		fmt.Printf("index: %d , value: %d\n", index, number)
		sum += number
	}
	return sum
}

func main() {
	fmt.Println(functions.Multiply(2, 2))
	fmt.Println(functions.Multiply2(3, 4))
	name := "JohnMark"
	fmt.Println(functions.LenAndUpperOfString(name))

	nameLen, upperName := functions.LenAndUpperOfString(name)

	fmt.Printf("My Name is %s , Name Length is %d\n", upperName, nameLen)

	functions.RepeatMe("a", "b", "c", "d", "Zip")

	// naked return
	totalLen, underWord := functions.LengAndUnderOfString("JOHNMARK")
	fmt.Printf("%d %s\n", totalLen, underWord)

	// call other package function
	message := echo.EchoMessageToUpper("hello world")
	fmt.Println(message)

	// summerization all numbers
	fmt.Println(superAdd(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}
