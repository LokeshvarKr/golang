package main

/*
Q 1:  Print anagram pairs from the given list of strings.

An anagram of a string is another string that contains the same characters, only the order of characters can be different.
E.g listen & silent.

I/p:   ["yo", "act", "flop", "tac", "foo", "cat", "oy", "olfp"]

o/p:

(act,tac)
(act,cat)
(tac,cat)
(yo,oy)
(flop,olfp)


*/

import "fmt"

func checkAnagram(a string, b string) bool {

	a1 := []rune(a)
	b1 := []rune(b)
	m := make(map[rune]int)
	for _, c := range a1 {
		if _, ok := m[c]; ok {
			m[c] = m[c] + 1
		} else {
			m[c] = 1
		}
	}

	for _, c := range b1 {
		v, ok := m[c]
		if !ok || v == 0 {
			return false
		} else {
			m[c] = m[c] - 1
		}
	}

	for _, v := range m {
		if v > 0 {
			return false
		}
	}

	return true
}

func main() {
	input := []string{"yo", "act", "flop", "tac", "foo", "cat", "oy", "olfp"}
	n := len(input)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if checkAnagram(input[i], input[j]) {
				fmt.Println("(", input[i], ",", input[j], ")")
			}
		}
	}

}
