package main

import (
	"fmt"
)

/*
//array declaration
func main() {
	var arr1 = [5]int{1, 2, 3, 4, 5}
	fmt.Println("array :", arr1)
}
*/

/*
//short hand declaration
func main() {
	arr1 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("array :", arr1)
}
*/

/*
//array iteration
func main() {
	n := 5
	var arr1 = [5]int{1, 2, 3, 4, 5}
	fmt.Println("array :", arr1)
	for i := 0; i < n; i++ {
		fmt.Println(arr1[i])
	}
}
*/

/*
func main() {
	arr1 := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("Type :%T ", arr1)

}
*/

/*
func main() {
	arr1 := [3][4]int{{1, 2, 3}, {6, 8, 9, 7}}
	fmt.Println("array :", arr1)
	fmt.Printf("Type :%T ", arr1)
}
*/

/*
//slices
func main() {
	items := []int{1, 2, 3}
	fmt.Println("Items :", items)
*/

func main() {
	arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("arr : ", arr)
	slice1 := arr[0:6]
	fmt.Println("slice1 : ", slice1)
	slice2 := arr[0:]
	fmt.Println("slice2 : ", slice2)
	slice3 := arr[:6]
	fmt.Println("slice3 : ", slice3)
	slice4 := arr[:]
	fmt.Println("slice4 : ", slice4)
}
