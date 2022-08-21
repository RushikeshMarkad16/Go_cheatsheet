package main

import (
	"fmt"
)

//struct
func main() {
	type employee struct {
		firstname string
		lastname  string
		age       int
		roll      int
	}

	emp := employee{
		firstname: "Stuart",
		lastname:  "Anderson",
		age:       29,
	}

	fmt.Printf("Employee name : %v %v \n", emp.firstname, emp.lastname)
	emp.firstname = "James"
	fmt.Printf("Employee name : %v %v %v \n", emp.firstname, emp.lastname, emp.roll)
}
