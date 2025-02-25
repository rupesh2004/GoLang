package main

import (
	"fmt"
	"sync"
)


func addfunction(chn1 chan int, wg *sync.WaitGroup, receivedData chan int){
	defer wg.Done()
	res:=0
	for val := range chn1{
		res+=val
	}
	receivedData<-res
	close(receivedData)

}
func main() {
	fmt.Println("Starting channel")

	var wg sync.WaitGroup
	wg.Add(1)
	chn1 := make(chan int,2)
	receiveData := make(chan int, 1)

	go addfunction(chn1, &wg,receiveData)
	chn1 <- 10
	chn1 <- 20
	close(chn1)

	result := <- receiveData
	fmt.Println("addition : ",result)
	wg.Wait()
	fmt.Println("end of main method")

}