package otherImportant

import (
	"math"
)

type myFloat float64

func (f myFloat) Ceil() float64 {
	return math.Ceil(float64(f))
}

func (f myFloat) Floor() float64 {
	return math.Floor(float64(f))
}
