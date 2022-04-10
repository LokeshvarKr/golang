package main

import (
	"fmt"
)

func square(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i * i
	}
	close(ch)
}
func main() {
	ch := make(chan int, 5)
	go square(ch)
	for {
		fmt.Println("len: ", len(ch))
		val, ok := <-ch
		if ok == false {
			fmt.Println("Loop Broke")
			break
		} else {
			fmt.Println(val, ok, len(ch), cap(ch))
		}
	}
}
