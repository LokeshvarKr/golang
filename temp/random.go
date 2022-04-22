package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Q. convert a large random integer into base 58

*/

const MIN int = 1000000000
const MAX int = 9999999999

// Not considering 0(zero), capital O , 1 (one) and small l
var LookUP []rune = []rune("abcdefghijkmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ23456789")

func base58(t int) string {
	output := ""
	for t > 0 {
		x := t % 58
		t = t / 58
		output += string(LookUP[x])
	}

	temp := []rune(output)
	// fmt.Println(string(temp))
	for i, j := 0, len(temp)-1; i < j; i, j = i+1, j-1 {
		temp[j], temp[i] = temp[i], temp[j]
	}
	return output
}

var store = make(map[int]string)

func main() {

	t := time.Now().UnixNano()
	rand.Seed(t)

	repeat := 0
	for count := 0; count < 1000; count++ {
		tempInt := rand.Intn(MAX-MIN+1) + MIN
		if _, ok := store[tempInt]; ok {
			repeat++
			continue
		}
		output := base58(tempInt)
		store[tempInt] = output
		// fmt.Println(output)
	}

	fmt.Println(repeat, len(store))

}
