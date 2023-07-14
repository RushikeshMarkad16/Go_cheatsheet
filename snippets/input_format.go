package main

import (
	//"bufio"
	"fmt"
	//"os"
)

func main() {
	fmt.Println("Hello World")
	var str string
	fmt.Scanln(&str)   //for single word
	fmt.Println(str)
// 	reader := bufio.NewReader(os.Stdin)    //for sentence
// 	input, err := reader.ReadString('\n')
// 	fmt.Println(input)
// 	fmt.Println(err)
}
