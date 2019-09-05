package main

import (
	"fmt"
)

func main() {
	/*
	* Write a program to print out
	*	G
	*	GG
	*	GGG
	* GGGG
	*	...
	 */
	// for i := 0; i < 25; i++ {
	// 	for j := 0; j <= i; j++ {
	// 		fmt.Print("G")
	// 	}
	// 	fmt.Print("\n")
	// }

	/*
	* Write a program to print out
	*	GGGGGGGGGGGGGGG
	*	G             G
	*	G             G
	*	G             G
	*	G             G
	*	GGGGGGGGGGGGGGG
	 */
	// for i := 0; i < 10; i++ {
	// 	for j := 0; j <= 23; j++ {
	// 		if i == 0 || i == 9 || j == 0 || j == 23 {
	// 			fmt.Print("G")
	// 		} else {
	// 			fmt.Print(" ")
	// 		}
	// 	}
	// 	fmt.Print("\n")
	// }
	/*
	* Write a program to print out
	*	G
	*	GG
	*	GGG
	*	GGGG
	*	GGGGG
	*	GGGG
	*	GGG
	*	GG
	*	G
	 */
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i > j {
				fmt.Print("G")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}
