// Merging output of two channels
// use of nil channel

package main

import (
	"fmt"
	"time"
)

func portal11(ch *chan int) {
	for i := 1; i <= 10; i++ {
		time.Sleep(time.Second * 1)
		*ch <- i
	}
	close(*ch)
}
func portal22(ch *chan int) {
	for i := 1; i < 20; i++ {
		time.Sleep(time.Second * 1)
		*ch <- (i * i)
	}
	close(*ch)
}

func main() {
	R1 := make(chan int)
	R2 := make(chan int)
	R3 := make(chan int)
	go portal11(&R1)
	go portal22(&R2)

	go func() {
		defer close(R3)
		var sum int = 0
		for R1 != nil || R2 != nil {
			select {
			case v, ok := <-R1:
				if !ok {
					fmt.Println("R1 is done")
					R1 = nil
				} else {
					fmt.Println("R1: ", v)
					sum += v
					R3 <- sum
				}
			case v, ok := <-R2:
				if !ok {
					fmt.Println("R2 is done")
					R2 = nil
				} else {
					fmt.Println("R2: ", v)
					sum += v
					R3 <- sum
				}
			}
		}
	}()

	for v := range R3 {
		fmt.Println("R3: ", v)
	}
	fmt.Println("R3 is done")
}
