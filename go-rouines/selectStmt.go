package main

import (
	"fmt"
	"sync"
)

func fun1(chn1 chan string, wg *sync.WaitGroup){
	chn1<-"hello from func1"
	defer wg.Done()
}
func fun2(chn2 chan string, wg *sync.WaitGroup){
	chn2<-"hello from func2"
	defer wg.Done()

}

func main() {
	fmt.Println("main started")
	var wg sync.WaitGroup
	wg.Add(2)

	chn1:= make(chan string,1)
	chn2:= make(chan string,1)
	
	go fun1(chn1,&wg)
	go fun2(chn2,&wg)

	select{
		case msg1 := <-chn1:
			fmt.Println(msg1)
		case msg2 := <-chn2:
			fmt.Println(msg2)
	}

	wg.Wait()
	fmt.Println("main ended")
}
