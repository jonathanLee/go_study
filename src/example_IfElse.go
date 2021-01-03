package main

import "fmt"

func main() {

	var i = 0
	var j = 1

	//block 을 아래와 같이 해야함.
	if i > j {
		fmt.Println("i is largest")
	} else {
		fmt.Println("j is largest")
	}
}
