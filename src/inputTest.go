package main

import "fmt"

func main() {

	/*
		입력값을 처리할 때는 Scanln
	*/

	fmt.Println("당신의 이름은?")
	name := ""
	fmt.Scanln(&name)

	fmt.Println("내 이름은 " + name)
}
