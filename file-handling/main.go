package main

import (
	"fmt"
	"io"
	"os"
)



func main() {
	fmt.Println("File handling demo")
	// Create file
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	fmt.Println("File created:", file.Name())

	// Write to the file
	_, err = file.WriteString("Hello from Rupesh Bhosale")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("File written successfully")

	file, err = os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close() 

	// Read the file
	fmt.Println("Reading file contents:")
	buffer := make([]byte, 1024)
	for {
		n, err := file.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break // End of file
			}
			fmt.Println("Error reading file:", err)
			return
		}
		fmt.Println(string(buffer[:n]))
	}

	// read file using read file
	content,err := os.ReadFile(file.Name())
	if err!=nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(content))

	// delete file
	_ = os.Remove(file.Name())
	fmt.Println("fie deleted")

	// close connection
	file.Close()

	

}
