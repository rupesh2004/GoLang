package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader:=bufio.NewReader(os.Stdin)

	fmt.Println("Enter value A : ")
	input,_ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	num1,err := strconv.Atoi(input)

	fmt.Println("Enter value B : ")
	input2,_ := reader.ReadString('\n')
	input2 = strings.TrimSpace(input2)
	num2,err :=strconv.Atoi(input2)

	if err != nil{
		fmt.Println("Invalid input")
	}
	
	if num1 > num2 {
		fmt.Println("A is greater than B")
	}else{
		fmt.Println("B is greater than A")
	}
}