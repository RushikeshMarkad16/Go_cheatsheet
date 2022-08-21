package main

import "fmt"

/*
//for loop
func main() {
	sum := 0
	for i := 1; i < 5; i++ {
		sum = sum + i
	}

	fmt.Println("Sum is :", sum)
}
*/

/*
while loop
func main() {
	n := 1
	for n < 5 {
		n = n * 2
	}

	fmt.Println("n is :", n)
}
*/

/*
//ranging/slicing
func main() {
	strings := []string{"Hello", "World"}
	for i, s := range strings {
		fmt.Println(i, s)
	}
}
*/

/*
func main() {
	strings := []string{"Hello", "World"}
	for i := range strings {
		fmt.Println(i, strings[i])
	}
}
*/

//continue
func main() {
	n := 0
	for i := 1; i < 10; i++ {
		n = n + 1
		if n == 5 {
			continue
		}
		fmt.Println("n is :", n)

	}
}

/*
//break
func main() {
	n := 1
	for i := 1; i < 10; i++ {
		fmt.Println("n is :", n)
		n = n + 1

		if n == 5 {
			break
		}
	}
}
*/
