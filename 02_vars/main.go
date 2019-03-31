package main

import "fmt"

func main() {
	// Main Types
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// Using var
	// var name = "Chisimdi"
	// name := "Simdi"
	var age int32 = 37
	const isCool = true
	var size = 1.3

	name, email := "Damian", "simdi@yahoo.com"

	fmt.Println(name, email, age, isCool, size)
	fmt.Printf("%T\n", size)
}
