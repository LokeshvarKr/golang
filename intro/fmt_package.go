package main

import "fmt"

func main() {
	str1 := "This is str1"
	str2 := "This is str2"
	str3 := "This is str3"
	
	// fmt.Println(str1,str2,str3)
	
	// stringLength,err := fmt.Println(str1,str2,str3)
	// if err == nil {
	//	fmt.Println(stringLength)
	//}
	
	
	stringLength, _ := fmt.Println(str1,str2,str3)
	fmt.Println("printed String Length:", stringLength)

	aNumber := 42
	isTrue := true

	fmt.Printf("Value of aNumber: %v\n",aNumber)
	fmt.Printf("Value of aNumber: %v\n",isTrue)
	fmt.Printf("Value of aNumber as float: %.2f\n",float64(aNumber))
	fmt.Printf("Data Types : %T, %T, %T, %T and %T\n",str1,str2,str3,aNumber,isTrue)
	returned_string := fmt.Sprintf("Data Types : %T, %T, %T, %T and %T\n",str1,str2,str3,aNumber,isTrue)
	fmt.Println(returned_string)


}