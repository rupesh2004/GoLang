package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{Name: "John", Age: 30}
	fmt.Println("Object Data : ", p)
	// Marshal the struct to JSON
	jsonData, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("JSON Data : ", string(jsonData))
	// unmarshal json data
	var objectData Person
	err = json.Unmarshal(jsonData,&objectData)
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Println("Object data : ",objectData)

	
}
