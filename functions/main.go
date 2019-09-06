package main

import (
	"fmt"
)

// func (st *Stack) Pop() int {
// 	v := 0
// 	for ix := len(st) - 1; ix >= 0; ix-- {
// 		if v = st[ix]; v != 0 {
// 			st[ix] = 0
// 			return v
// 		}
// 	}
// }

// Multipy numbers passed as parameters
func MultiPly3Nums(a int, b int, c int) int {
	// var product int = a * b * c
	// return product
	return a * b * c
}

func main() {
	// Pop()
	fmt.Printf("Multiply 2 * 5 * 6 = %d\n", MultiPly3Nums(2, 5, 6))
	var i1 int = MultiPly3Nums(3, 10, 8)
	fmt.Printf("Multiply 3 * 10 * 8 = %d\n", i1)
}
