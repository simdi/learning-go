package main

import (
	"fmt"
)

func main() {
	sample := "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	fmt.Printf("sample: %s, Type: %T\n", sample, sample)

	// Looping through the string with for loop
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x \n", sample[i])
	}

	// Looping through the string with for range
	for i, v := range sample {
		fmt.Printf("%d, %x \n", i, v)
	}
}
