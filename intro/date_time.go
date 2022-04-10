package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	time.Sleep(10*time.Second)
	t := time.Now()
	fmt.Println(start," , ",t)

	fmt.Println("Difference: ",t.Sub(start))

	t1 := time.Date(2009,time.November,10,23,0,0,0,time.UTC)
	fmt.Printf("Go Launched at : %s\n",t1)

	tomorrow := t.AddDate(0,0,1)
	fmt.Println("Tomorrow: ",tomorrow)

}