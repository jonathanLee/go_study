/*
 * Go 언어는 세미콜론이 필요없다.
//+= 연산자는 s += sep + os.Args[i]
 * s = s + sep+ os.Args[i]
 * 변수 선언 방법
 * s := ""
*/
package main

import "fmt"
import "os"

func main() {
	//var s, sep string
	s, sep := "", ""
	//string 변수 s, sep 선언
	for i := 1; i < len(os.Args); i++ {
		fmt.Println(i)
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
