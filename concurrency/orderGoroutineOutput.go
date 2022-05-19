/*

Question : Print a b c in order 10 times 

a
b
c
a
b
c
a
b
c
...

*/


package main

import (
	"fmt"
	"time"
)
 
func printA(a chan int, b chan int) {
	for i := 0; i < 10; i++ {
		<-a
		fmt.Println("a")
		b <- 1
	}

}

func printB(b chan int, c chan int) {
	for i := 0; i < 10; i++ {
		<-b
		fmt.Println("b")
		c <- 1
	}
}

func printC(c chan int, a chan int) {
	for i := 0; i < 10; i++ {
		<-c
		fmt.Println("c")
		a <- 1
	}
}

func main() {
	a, b, c := make(chan int), make(chan int), make(chan int)
	go printA(a, b)
	go printB(b, c)
	go printC(c, a)
	a <- 1
	time.Sleep(time.Second * 2)
	<-a
}
