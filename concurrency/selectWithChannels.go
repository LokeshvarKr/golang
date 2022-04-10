package main

import (
	"fmt"
	"time"
)

func portal11(ch chan string) {
	ch <- "Welcome to channel 1"
}
func portal22(ch chan string) {
	ch <- "Welcome to channel 2"
}

func main() {
	R1 := make(chan string)
	R2 := make(chan string)

	go portal11(R1)
	go portal22(R2)

	time.Sleep(1 * time.Second)

	select {
	case op1 := <-R1:
		fmt.Println(op1)
	case op2 := <-R2:
		fmt.Println(op2)
	}
}
