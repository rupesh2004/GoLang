package main

import "fmt"

func simpleFunction() {
	fmt.Println("This is simple function")
}
func add(a, b int) int {
	return a+b
}
func multiply(a,b int)(result int){
	result=a*b
	return
}
func multipleReturn() (int,int) {
	return 3, 4
}
func variadicFunction(nums ...int){
	fmt.Println(nums)
	tot:=0
	for _,num := range nums {
		tot+=num
	}
	fmt.Println(tot)
	fmt.Printf("%T",nums)
}
func main() {
	simpleFunction()
	res:=add(3,4)
	fmt.Println("Add : ",res)
	fmt.Println("Multiplication : ",multiply(3,4))
	a,b := multipleReturn()
	fmt.Println("Multiple Return : ",a,b)
	variadicFunction(1,2,3,4,5)
}