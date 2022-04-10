package main

import "fmt"

// Mulitple inheritance 


type IBase1 interface{
	Base1Method()
}

type IBase2 interface{
	Base2Method()
}

type Base1 struct{
	Property1 string
}
 
func(b Base1) Base1Method(){
	fmt.Println("From Base1Method, Property1:",b.Property1)
}


type Base2 struct{
	Property2 string
}

func(b Base2) Base2Method(){
	fmt.Println("From Base2Method, Property1:",b.Property2)
}

type Child struct{
	Base1
	Base2
	ChildProperty string
}

func(c Child) ChildMthod(){
	fmt.Println("From ChildMethod, Child1Property:",c.ChildProperty)
	c.Base1Method()
	c.Base2Method()
}

func Check1(b IBase1){
	b.Base1Method()
}

func Check2(b IBase2){
	b.Base2Method()
}

func main(){

	b1 := Base1{
		Property1: "Property1",
	}
	b2 := Base2{
		Property2: "Property2",
	}

	c := Child{
		Base1: b1,
		Base2: b2,
		ChildProperty: "ChildProperty",
	}

	c.ChildMthod()
	Check1(c)
	Check2(c)

}


