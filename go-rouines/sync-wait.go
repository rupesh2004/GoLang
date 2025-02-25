package main

import (
	"fmt"
	"sync"
)

func worker(id int , wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d started\n", id)
	fmt.Printf("Worker %d ended\n", id)
}

func main() {
	var wg sync.WaitGroup
	fmt.Println("worked started")
	for i:=0;i<3;i++{
		wg.Add(1)
		go worker(i,&wg)
	}

	wg.Wait()
	fmt.Println("worked ended")


}
