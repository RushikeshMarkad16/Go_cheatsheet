package main

import (
	"fmt"
	"sync"
)
//WaitGroups
var wg sync.WaitGroup

func main() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println("Hello World")
}
