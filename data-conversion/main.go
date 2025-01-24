package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Datatype conversion")
	str1:="1234"

	//str to int
	num1,_ :=strconv.Atoi(str1)
	fmt.Println("Value of num1:",num1)
	fmt.Printf("Type of num1 : %T",num1)

	//int to str
	str2 := strconv.Itoa(num1)
	fmt.Println("\nValue of str1:",str2)
	fmt.Printf("Type of str1 : %T",str2)

	//str to float
	str3 :="3.14"
	num3,_ := strconv.ParseFloat(str3,64)
	fmt.Println("\nValue of num3:",num3)
	fmt.Printf("Type of num3 : %T",num3)

	// float to int
	num4 := int(num3)
	fmt.Println("\nValue of num4:",num4)
	fmt.Printf("Type of num4 : %T",num4)


}