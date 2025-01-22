package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// var name string
	// fmt.Println("Enter your name : ")
	// fmt.Scanln(&name)
	// fmt.Println("Hello,", name)

	reader := bufio.NewReader(os.Stdin)
	username, _ := reader.ReadString('\n')
	fmt.Println("Hello", username)	

	

}	
