package main

import "fmt"

func main() {

	type Car struct {
		Name  string
		Price float32
	}

	car1 := Car{
		Name:  "Tesla",
		Price: 100,
	}

	fmt.Println(car1.Name)
	fmt.Println(car1.Price)
}
