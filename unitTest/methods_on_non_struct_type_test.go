package otherImportant

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCeil(t *testing.T) {

	var f myFloat = 34.5
	var expectedCeil float64 = 35
	var expectedFloor float64 = 34

	ceil := f.Ceil()
	floor := f.Floor()
	assert.Equal(t, expectedCeil, ceil, "Error")
	assert.Equal(t, expectedFloor, floor, "Error")

}
