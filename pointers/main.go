package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func changeTheValueOfAPointer(value *int) {
	*value = 300
	fmt.Printf("This pointer is of type %T\n", value)
	return
}

// Indirectly modify parameter.
// This is not ideal.
func changeValueLocally(val int) int {
	fmt.Printf("The original value is: %d\n", val)
	b := &val
	*b = 6
	return val
}

func main() {
	b := 100
	a := &b
	var c *int
	c = &b
	// e := "my string"
	// f := &e
	fmt.Println("a", a)
	fmt.Println("b", b)
	fmt.Println("c", c)
	fmt.Printf("a %T\n", a)
	fmt.Printf("b %T\n", b)
	fmt.Printf("c %T\n", c)
	changeTheValueOfAPointer(c)
	fmt.Printf("a pointer %p\n", a)
	fmt.Printf("c pointer %p\n", c)
	fmt.Printf("b pointer %v\n", b)
	fmt.Printf("b pointer %v\n", *a)
	fmt.Printf("%T\n", os.Args)
	for _, v := range os.Args {
		fmt.Println("Values", v)
	}

	fmt.Println("usage", filepath.Base(os.Args[0]))

	val := 5
	fmt.Printf("Changed value is: %d\n", changeValueLocally(val))
}
