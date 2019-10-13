package main

import "fmt"

type Option struct {
	num     int
	name    string
	balance float64
}

func (st *Stack) Pop() int {
	v := 0
	for ix := len(st) - 1; ix >= 0; ix-- {
		if v = st[ix]; v != 0 {
			st[ix] = 0
			return v
		}
	}
}

// Multipy numbers passed as parameters
func MultiPly3Nums(a int, b int, c int) int {
	// var product int = a * b * c
	// return product
	return a * b * c
}

func variadicFunction(a int, b int, c ...int) (a1 int, b1 int, c1 []int) {
	return a, b, c
}

func variadicFuncWithOptionType(a int, b int, c Option{}) (a1 int, b1 int, c1 Option{}) {
	fmt.Printf("a: %v", a)
	return a, b, c
}

func a() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

func f() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d\n", i)
	}
}

func main() {
	Pop()
	fmt.Printf("Multiply 2 * 5 * 6 = %d\n", MultiPly3Nums(2, 5, 6))
	var i1 int = MultiPly3Nums(3, 10, 8)
	fmt.Printf("Multiply 3 * 10 * 8 = %d\n", i1)

	a, b, c, d := 2, 54, 64, 453
	e := Option{name: "John Doe"}
	r1, r2, r3 := variadicFunction(a, b, c, d)
	v1, v2, v3 := variadicFuncWithOptionType(a, b, e)
	a()
	f()
	fmt.Printf("Variadic Values: a = %d, b = %d, c = %v\n", r1, r2, r3)
	fmt.Printf("Variadic option Values: a = %d, b = %d, c = %v\n", v1, v2, v3)
}
