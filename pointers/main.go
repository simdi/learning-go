package main

import (
	"fmt"
)

func main() {
	b := 100
	a := &b
	fmt.Println("a", a)
	fmt.Println("b", b)
	fmt.Printf("a %T\n", a)
	fmt.Printf("b %T\n", b)
}
