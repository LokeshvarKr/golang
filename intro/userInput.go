package main

import (
	"fmt"
)

func main() {
	var intVal int
	var stringVal string
	fmt.Println("Enter stringVal : ")
	fmt.Scanf("%s", &stringVal)
	fmt.Println("Enter intVal : ")
	fmt.Scanf("%d", &intVal)
	
	fmt.Printf("intVal : %d and stringVal : %s", intVal, stringVal)
}
