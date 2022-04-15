// Tree Constructor
// Have the function TreeConstructor(strArr) take the array of strings stored in strArr,
// which will contain pairs of integers in the following format: (i1,i2),
// where i1 represents a child node in a tree and the second integer i2 signifies that it is the parent of i1.
// For example: if strArr is ["(1,2)", "(2,4)", "(7,2)"], then this forms the following tree:
//       4
//      /
//     2
//    /  \
//   1     7

// which you can see forms a proper binary tree. Your program should,
// in this case, return the string true because a valid binary tree can be formed.
// If a proper binary tree cannot be formed with the integer pairs, then return the string false.
// All of the integers within the tree will be unique, which means there can only be one node in the tree with the given integer value.
// Examples
//Input: []string {"(1,2)", "(2,4)", "(5,7)", "(7,2)", "(9,5)"}
//Output: true
//Input: []string {"(1,2)", "(3,2)", "(2,12)", "(5,2)"}
//Output: false
// Tags: Google Facebook

package main

import (
	"fmt"
	"strings"
)

func TreeConstructor(strArr []string) string {
	parentChildren := make(map[string][]string)
	childParent := make(map[string]string)
	for _, e := range strArr {
		e = strings.Trim(e, "()")
		pair := strings.Split(e, ",")
		parentChildren[pair[1]] = append(parentChildren[pair[1]], pair[0])
		if len(parentChildren[pair[1]]) > 2 {
			return "false"
		}
		childParent[pair[0]] = pair[1]
	}

	//check if only one root is present
	root := 0
	for parent := range parentChildren {
		_, ok := childParent[parent]
		if !ok {
			root += 1
		}
		if root > 1 {
			return "false"
		}
	}
	return "true"
}

func main() {

	var strArr []string = []string{"(1,2)", "(2,4)", "(5,7)", "(7,2)", "(9,5)"}
	fmt.Println(TreeConstructor(strArr))

}
