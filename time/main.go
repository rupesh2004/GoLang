package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time demo")
	fmt.Println("Current Time : ",time.Now())
	fmt.Println("Formatted current time : ",time.Now().Format("02-01-2006, Monday"))
	fmt.Println("Formatted current time : ",time.Now().Format("03:04PM"))

	t,_:=time.Parse("2006/01/02","2025/01/26")
	fmt.Println(t)

	fmt.Println("Later day:",time.Now().Add(24*time.Hour).Format("2006/01/02, Monday"))
	fmt.Println("Previous days:",time.Now().Add(-24*time.Hour).Format("2006/01/02, Monday"))

	//sleep
	fmt.Println("stoping execution fot 2 seconds")
	time.Sleep(10*time.Second)
	fmt.Println("2 seconds passed")


}
