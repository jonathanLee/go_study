package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	//현재 경로를 보여주는 함수
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("디렉토리 경로 : " + dir)

}
