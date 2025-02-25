package main
import "fmt"
func sending(s chan<- string) {
	s <- "rupesh"
}
func main() {

	mychanl := make(chan string)
	go sending(mychanl)
	fmt.Println(<-mychanl)
}
