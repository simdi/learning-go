package main

import (
	"fmt"
)

func Add2() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func Adder(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}

func returnDefer() (ret int) {
	defer func() {
		ret++
	}()
	fmt.Printf("Value %v\n", ret)
	return ret
}

func main() {
	fmt.Println("this is what it prints", returnDefer())
	// make an Add2 function, give it a name p2, and call it:
	p2 := Add2()
	fmt.Printf("Call Add2 for 3 gives: %v\n", p2(3))
	// make a special Adder function, a gets value 3:
	TwoAdder := Adder(2)
	fmt.Printf("The result is: %v\n", TwoAdder(3))
}
