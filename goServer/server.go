package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from rupesh")
		fmt.Fprint(w,"\nhello")
		
	})

	fmt.Println("starting server with port 8080")

	error :=http.ListenAndServe(":8080",nil)
	if error != nil {
		fmt.Println("error starting server")
		return
	}

}