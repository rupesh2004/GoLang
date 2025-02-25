package main

import (
	"fmt"
	"time"
)

func add1(){
	fmt.Println("in add 1 function")
	fmt.Println(7+4)
}
func add2(){
	fmt.Println("in add 2 function")
	fmt.Println(4+4)
}

func main() {
	fmt.Println("Hello")
	go add1()
	go add2()

	//time.Sleep(1 * time.Second)
	time.Sleep(1*time.Millisecond)
}