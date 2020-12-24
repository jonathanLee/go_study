package main

import "fmt"

/*
* @author : hoon
* @Description : 포인터는 값 자체보다는 메모리상의 위치를 가리킨다.
 * 는 포인트 변수를 선언할때 사용
 & 는 포인트의 메모리상의 주소값을 가져온다.
*/
func main() {
	var x *int
	var num = 12
	x = &num

	fmt.Println("num : ", num)
	fmt.Println("num pointer", x)

}
