package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Todo struct {
	UserId    int `json:"userId"`
	Id        int `json:"id"`
	Title     string `json:"title"`
	Completed bool `json:"completed"`
}

func main() {
	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()
	// data, err := ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(string(data))

	var todo Todo
	err = json.NewDecoder(response.Body).Decode(&todo)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(todo)

}
