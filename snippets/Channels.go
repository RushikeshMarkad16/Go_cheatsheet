/*
//Unbuffered Channels
package main

import (
	"fmt"
	"time"
)

func DoSm(s string, ch chan int) {
	ch <- 8
	fmt.Println(s)
}

func DoSm1(s string, ch chan int) {
	l := <-ch
	fmt.Println(l)
	fmt.Println(s)
}

func main() {
	ch := make(chan int)
	go DoSm("World", ch)
	go DoSm1("Hello", ch)
	time.Sleep(1 * time.Second)

}
*/

//Buffered Channel (has length argument)
package main

import (
	"fmt"
	"time"
)

func DoSm(s string, ch chan int) {
	ch <- 8
	fmt.Println(s)
}

func DoSm1(s string, ch chan int) {
	//l := <-ch
	//fmt.Println(l)
	fmt.Println(s)
}

func main() {
	ch := make(chan int, 1)
	go DoSm("World", ch)
	go DoSm1("Hello", ch)
	time.Sleep(1 * time.Second)

}
