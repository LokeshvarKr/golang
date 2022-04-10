package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	i1, i2, i3 := 10, 11, 12
	intSum := i1 + i2 + i3
	fmt.Println("intSum: ", intSum)

	f1, f2, f3 := 10.1, 11.2, 12.3
	floatSum := f1 + f2 + f3
	fmt.Println("floatSum: ", floatSum)

	var b1, b2, b3, bigSum big.Float
	b1.SetFloat64(10.1)
	b2.SetFloat64(11.2)
	b3.SetFloat64(12.3)

	bigSum.Add(&b1, &b2).Add(&bigSum, &b3)
	// fmt.Printf("bigSum: %.10g\n",bigSum)
	fmt.Printf("bigSum: %.10g\n", &bigSum)

	circleRadius := 2.3
	circleCircumference := circleRadius * math.Pi
	fmt.Printf("circleCircumference: %.2f", circleCircumference)

}
