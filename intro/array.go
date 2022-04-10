package main

import "fmt"


type Student struct{
	x int
	y int
}

func main(){
	arr1:= [5]int{}
	for i,j := range arr1{
		fmt.Println(i,j);
	}

	fmt.Printf("%T",arr1)

	arr2:=[]int{}
	fmt.Printf("%T",arr2)

}


