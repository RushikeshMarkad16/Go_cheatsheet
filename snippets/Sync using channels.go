//Sync and Mutex
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// var x = 0

// func increment(w *sync.WaitGroup, m *sync.Mutex) {
// 	m.Lock()
// 	x = x + 1
// 	m.Unlock()
// 	w.Done()
// }

// func main() {
// 	var w sync.WaitGroup
// 	var m sync.Mutex
// 	for i := 0; i < 1000; i++ {
// 		w.Add(1)
// 		go increment(&w, &m)
// 	}
// 	w.Wait()
// 	fmt.Println("x is :", x)
// }


//Using Channels
package main

import (
	"fmt"
	"sync"
)

var x = 0

func increment(w *sync.WaitGroup, ch chan bool) {
	ch <- true
	x = x + 1
	<- ch
	w.Done()
}

func main() {
	var w sync.WaitGroup
	ch := make(chan bool)
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w, ch)
	}
	w.Wait()
	fmt.Println("x is :", x)
}
