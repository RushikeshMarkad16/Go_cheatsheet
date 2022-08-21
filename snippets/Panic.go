package main

import (
	"fmt"
)

/*
func main() {
	arr := []int{0, 1, 2, 3}
	fmt.Println("10th elelment of slice : ", arr[10])
	fmt.Println("Normal execution done")
}
*/

/*
func FirstName(name string) {
	if name == "" {
		panic("Runtime error : Name cannot be empty")
	}
}

func main() {
	fmt.Println("Before function")
	FirstName("")
	fmt.Println("After function")
*/

func FirstName(name string) {
	defer fmt.Println("This is defer statement")
	if name == "" {
		panic("Runtime error : Name cannot be empty")
	}
}

func main() {
	fmt.Println("Before function")
	defer fmt.Println("Defer in int")
	FirstName("")
	defer fmt.Println("Defer after function")
	fmt.Println("After function")
}
