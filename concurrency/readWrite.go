package main

import (
	"fmt"
	"time"
)

func generateNumber(ch chan<- int) {
	for i := 1; i < 10000; i++ {
		ch <- i
	}
	close(ch)
}

func calculateSquare(ch <-chan int) {
	// for {
	// 	i, ok := <-ch
	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Println(i, i*i)
	// }

	for i := range ch {
		fmt.Println(i, i*i)
	}
}

func main() {
	ch := make(chan int)
	go generateNumber(ch)
	go calculateSquare(ch)
	time.Sleep(time.Second)
}
