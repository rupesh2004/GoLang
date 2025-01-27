package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func fetchData() {
	fmt.Println("Fetching data...")
	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()

	var todo Todo
	err = json.NewDecoder(response.Body).Decode(&todo)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Fetched Data:", todo)
}

func postData() {
	fmt.Println("Posting data...")
	todo := Todo{
		UserId:    2004,
		Id:        201,
		Title:     "data uploading using post",
		Completed: false,
	}
	// converting to json
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println(err)
		return
	}
	// convert json to reader
	reader := bytes.NewReader(jsonData)
	res, err := http.Post("https://jsonplaceholder.typicode.com/todos/", "application/json", reader)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Response Data:", string(data))
}

func updateData() {
    fmt.Println("Updating data...")
    todo := Todo{
        UserId:    2004,
        Id:        201,
        Title:     "data updating using put",
        Completed: true,
    }
    // Convert to JSON
    jsonData, err := json.Marshal(todo)
    if err != nil {
        fmt.Println("Error marshaling data:", err)
        return
    }

    // Send PUT request
    reader := bytes.NewReader(jsonData)
    req, err := http.NewRequest("PATCH", "https://jsonplaceholder.typicode.com/todos/201", reader)
    if err != nil {
        fmt.Println("Error creating request:", err)
        return
    }
    req.Header.Set("Content-Type", "application/json")

    res, err := http.DefaultClient.Do(req)
    if err != nil {
        fmt.Println("Error sending request:", err)
        return
    }
    defer res.Body.Close()

    if res.StatusCode != http.StatusOK {
        fmt.Printf("Server responded with status: %d\n", res.StatusCode)
        return
    }

    data, _ := ioutil.ReadAll(res.Body)
    fmt.Println("Updated Data:", string(data))
}

func deleteData(){
	fmt.Println("Deleting data...")
	// Send DELETE request	
	req, err := http.NewRequest("DELETE", "https://jsonplaceholder.typicode.com/todos/201",nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	res,err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer res.Body.Close()
	fmt.Println("Status code : ",res.Status)
}


func main() {
	postData()
	updateData()
	deleteData()
	fetchData()
}
