package main

import "fmt"

func main() {
	fmt.Println("Maps demo")
	map1 := make(map[string]int)
	map1["one"] = 1
	map1["two"] = 2
	fmt.Println(map1)
	fmt.Println("One: ",map1["one"])
	delete(map1,"two")
	
	num,exist:=map1["two"]
	fmt.Println("Two: ",num,"exist: ",exist)

	num,exist= map1["one"]
	fmt.Println("One: ",num,"exist: ",exist)
	
	map1["three"] = 3
	map1["four"]=4
	fmt.Println(map1)

	for index,value :=range map1 {
		fmt.Println(index,value)
	}

	map1["one"] = 4
	fmt.Println(map1)

	clear(map1)
	fmt.Println(map1)

}