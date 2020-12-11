package functions

import (
	"fmt"
	"strings"
)

// return multiple value
func LenAndUpperOfString(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

// get all parameters
func RepeatMe(words ...string) {
	fmt.Println(words)
}

// naked return function
func LengAndUnderOfString(name string) (length int, lowerWord string) {
	defer fmt.Println("I'm Done")
	length = len(name)
	lowerWord = strings.ToLower(name)
	return
}
