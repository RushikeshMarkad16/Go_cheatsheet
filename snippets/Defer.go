package main

import (
	"fmt"
)

/*
func main() {
	fmt.Println("Hello World")
	fmt.Println("1")
	fmt.Println("2")
	fmt.Println("3")
	defer fmt.Println("4")
	fmt.Println("5")
	return
	fmt.Println("6")
}
*/

func main() {
	fmt.Println("Hello")

	for i := 0; i <= 3; i++ {
		defer fmt.Println(i)
	}
	
	fmt.Println("World")
}
