package main

import "fmt"

/*
function without arguments
func demo() {
	a := 10
	b := 16
	sum := a + b
	fmt.Println("Sum is :", sum)
}

func main() {
	demo()
}
*/

/*
function with arguments
func demo(a, b int) {
	sum := a + b
	fmt.Println("Sum is :", sum)
}

func main() {
	demo(5, 8)
}
*/

/*
func demo(a, b int) int {
	sum := a + b
	return sum
}

func main() {
	sum1 := demo(5, 8)
	fmt.Println("Sum is :", sum1)
}
*/

/*
func CalculateBill(price, no_of_items int) {
	total := price * no_of_items
	fmt.Println("Total bill is : ", total)
}

func main() {
	var price int = 50
	var no_of_items int = 6
	CalculateBill(price, no_of_items)
}
*/

func CalculateBill(price, no_of_items int) int {
	total := price * no_of_items
	return total
}

func main() {
	var price int = 50
	var no_of_items int = 6
	total := CalculateBill(price, no_of_items)
	fmt.Println("Total bill is : ", total)
}
