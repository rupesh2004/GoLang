package main

import "fmt"

func callByValue(a,b int){
	fmt.Println("Inside callByValue function")
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	// swap values
	a,b = b,a
	fmt.Println("Value of a =",a)
	fmt.Println("Value of b =",b)
}

func callByReference(a,b *int){
	fmt.Println("Inside callByReference function")
	fmt.Println("a =", *a)
	fmt.Println("b =", *b)
	// swap values
	*a, *b = *b, *a
	fmt.Println("Value of a =",*a)
	fmt.Println("Value of b =",*b)
}

func main() {
	fmt.Println("Pointers demo")
	
	num:=2
	ptr1:=&num
	fmt.Println("Value of num:",num)
	fmt.Println("Address of ptr1:",ptr1)
	fmt.Println("Value of ptr1:",*ptr1)

	a:=2
	b:=4
	callByValue(a,b)
	fmt.Println("Value of a:",a)
	fmt.Println("Value of b:",b)

	// call by reference
	callByReference(&a,&b)
	fmt.Println("Value of a:",a)
	fmt.Println("Value of b:",b)

}