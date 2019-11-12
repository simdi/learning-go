package main

import (
	"fmt"
)

func main() {
	// Define maps
	emails := make(map[string]string)
	// Assign key values
	emails["simdi"] = "simdi@yahoo.com"
	emails["john"] = "john@yahoo.com"

	fmt.Println(emails)
}
