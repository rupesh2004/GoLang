package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	cnt:=0
	for{
		fmt.Println("infinite loop")
		if cnt == 5 {
			break
		}
		cnt++
	}

	number := [...]int{1,2,3,4}
	for index,value :=range number {
		fmt.Println("index: ",index,"value: ",value)
	}

}