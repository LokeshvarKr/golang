package main

import (
	"fmt"
)

func xyz(ch chan string) {
	<-ch
	<-ch
}
func main() {
	ch := make(chan string)
	go xyz(ch)
	ch <- "Lokesh"
	fmt.Println("one time")
	close(ch)
	ch <- "Nokia"  // this line will give error
	fmt.Println("second time")
}
