package main

import "fmt"



func main() {
	fmt.Println("Array demo")
	var name [5]string
	name[0] = "Rupesh"
	name[1] = "satya"

	fmt.Println(name[0],name[1])

	numbers := [5]int{1,2,3,4,5}
	fmt.Println(numbers)

	num := [...]int {1,2,3,4,6,7,54,4,4,6,7,4,5}
	fmt.Println(num)
	fmt.Println(len(num),len(numbers),len(name))

	// take array element from user
	var n[3] int
 	fmt.Println("Enter the array elements")
	for i := 0; i < 3; i++ {
		fmt.Println("Enter element ", i, " : ")
		fmt.Scanln(&n[i])
	}
	fmt.Println("You entered : ", n)
}
