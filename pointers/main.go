package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	b := 100
	a := &b
	fmt.Println("a", a)
	fmt.Println("b", b)
	fmt.Printf("a %T\n", a)
	fmt.Printf("b %T\n", b)
	fmt.Printf("%T\n", os.Args)
	for _, v := range os.Args {
		fmt.Println("Values", v)
	}

	fmt.Println("usage", filepath.Base(os.Args[0]))
}
