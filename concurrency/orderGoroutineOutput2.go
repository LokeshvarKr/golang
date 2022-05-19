package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup = sync.WaitGroup{}

func PrintA(current chan bool, next chan bool) {
	for i := 0; i < 10; i++ {
		<-current
		fmt.Println("a", i)
		next <- true
	}
	<-current
	wg.Done()
}

func PrintB(current chan bool, next chan bool) {
	for i := 0; i < 10; i++ {
		<-current
		fmt.Println("b", i)
		next <- true
	}
	wg.Done()
}
func PrintC(current chan bool, next chan bool) {
	for i := 0; i < 10; i++ {
		<-current
		fmt.Println("c", i)
		next <- true
	}
	wg.Done()
}

func main() {

	a, b, c := make(chan bool), make(chan bool), make(chan bool)
	wg.Add(3)
	go PrintA(a, b)
	go PrintB(b, c)
	go PrintC(c, a)

	a <- true
	wg.Wait()

}
