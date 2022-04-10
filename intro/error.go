package main

import (
	"fmt"
	"strings"
)

type MyCustomError struct {
	Code    string
	Message string
}

func (m MyCustomError) Error() string {
	return fmt.Sprintf("Error Code: %s , Reason: %s", m.Code, m.Message)
}

func checkName(name string) error {
	names := strings.Split(name, " ")
	if len(names) < 2 {
		return MyCustomError{Code: "NoSirName", Message: "Sir name is not present in the given name"}
	}
	return nil
}

func main() {
	err := checkName("Lokesh")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = checkName("Teegen Zongna")
	if err != nil {
		println(err.Error())
	}

}
