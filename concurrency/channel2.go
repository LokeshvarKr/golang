package main

import (
	"fmt"
)

func square(ch chan int) {
	for i := 0; i <= 10; i++ {
		ch <- i * i
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go square(ch)
	for {
		val, ok := <-ch
		if ok == false {
			fmt.Println(val, ok, " loop broke")
			break
		} else {
			fmt.Println(val, ok)
		}
	}
}
