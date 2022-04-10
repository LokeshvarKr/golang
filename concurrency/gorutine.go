package main

import (
	"fmt"
	"time"
)

func display(message string) {
	var i int
	for i = 0; i < 10; i++ {
		time.Sleep(2 * time.Second)
		fmt.Println(message, i)
	}
}
func main() {
	go display("Hey")
	// display("Hi")
	go display("Hello")
	display("byee")
	time.Sleep(20 * time.Second)
}
