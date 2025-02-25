package main

import "fmt"

func main() {
	var channel1 chan int
	fmt.Println("\nValue of channel : ",channel1)
	fmt.Printf("\nType of channel %T",channel1)

	// using cmake function
	channel2 :=make(chan int)
	fmt.Println("\nValue of channel : ",channel2)
	fmt.Printf("\nType of channel %T",channel2)
}