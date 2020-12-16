/*
 *
 */
package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println(os.Args[1])
}
