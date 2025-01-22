package main

import "fmt"

var PublicVariable = 345

func main() {
	var name string = "rupesh"
	fmt.Println(name)

	var value int = 10
	fmt.Println(value)

	var isMarried bool = true
	fmt.Println(isMarried)

	var decimalVal float64 = 8.4
	fmt.Println(decimalVal)

	const pi = 3.14
	fmt.Println(pi)

	var version = 4.5
	fmt.Println(version)

	var lname = "bhosale"
	fmt.Println(lname)

	PrintPublicVariable()
}
