package main

import "fmt"

func main() {
	// 명시적이 변수 선언
	var name string = "johnmark"
	fmt.Println(name)

	// 축약형 변수 선언, 타입추론을 통해 변수 자료형 선택
	message := "hi!"
	fmt.Println(message)
}
