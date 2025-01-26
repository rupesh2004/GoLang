package main

import "fmt"

func add(a,b int)(int){
	return a+b
}
func main() {
	
	fmt.Println("Differ keyword demo")
	defer fmt.Println("First line")
	defer fmt.Println("Second line")
	fmt.Println("Third line")

	res:=add(3,4)
	defer fmt.Println("Result:",res)

}