package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func doWork() int {
	time.Sleep(time.Second*1)
	return rand.Intn(1000) + 1
}

func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	go func() {
		for i := 1; i < 10000; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				x := doWork()
				ch <- x
			}()
		}

		wg.Wait()
		close(ch)
	}()

	for i := range ch {
		fmt.Println(i)
	}

}
