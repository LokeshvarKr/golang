package main

import "fmt"

// Type Hierarchy
// type inheritance is not allowed to go and hence it does not have type hierarchy.
// GO intentionally doesn’t allow this feature so any change in the behavior of an
// interface is only propagated to its immediate structures which defines all methods of the interface.

// Although we can implement type hierarchy using interfaces and struct like below

type IAnimal interface{
	Breathe()
}

type Animal struct {

}

func (a *Animal) Breathe() {
	fmt.Println("Animal breate")
}

type IAquatic interface {
	IAnimal
	Swim()
}

type Aquatic struct {
	Animal
}

func (a *Aquatic) Swim() {
	fmt.Println("Aquatic swim")
}

type INonAquatic interface{
	IAnimal
	Walk()
}

type NonAquatic struct {
	Animal
}

func (a *NonAquatic) Walk() {
	fmt.Println("Non-Aquatic walk")
}

type Shark struct {
	Aquatic
}

type Lion struct {
	NonAquatic
}

// The point to be noted here is that if you want distinction in your type 
// hierarchy where lets say a “shark”  should not be both “iAquatic” and “iNonAquatic”  ,
// then there should be at least one method in the method sets of “iAquatic” and “iNonAquatic”  
// which is not present in the other. In our example “swim”  and “walk”  are those methods.

// iAnimal
// ---iAquatic
// ------shark
// ---iNonAquatic
// ------lion

// Conclusion: Go doesn’t have support for type inheritance but the same can be achieved using
// embedding buts one needs to be careful while creating such kind of type hierarchy.
// Also, go does not provide method overriding.


func checkAquatic(a IAquatic) {

}

func checkNonAquatic(a INonAquatic) {

}

func checkAnimal(a IAnimal) {

}

func main() {
	shark := Shark{}
	lion := Lion{}
	shark.Breathe()
	shark.Swim()
	lion.Breathe()
	lion.Walk()
}