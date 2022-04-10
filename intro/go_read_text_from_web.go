package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	// "io/ioutil"
)

func main() {
	url := "https://golang.org/doc/"

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	fmt.Printf("Response Type : %T\n", resp)
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))

}
