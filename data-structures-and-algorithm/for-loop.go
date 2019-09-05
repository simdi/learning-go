package main

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
	// for i := 0; i < 10; i++ {
	// 	for j := 0; j < 10; j++ {
	// 		if i > j {
	// 			fmt.Print("G")
	// 		} else {
	// 			fmt.Print(" ")
	// 		}
	// 	}
	// 	fmt.Print("\n")
	// }

	/*
	* Try modifying a value in the for range loop
	* For range
	 */
	// names := []string{"John", "Doe", "James", "Bond"}
	// for i, v := range names {
	// 	// You can't change v in a for range.
	// 	v = "Changed Doe"
	// 	fmt.Printf("Index: %d, Value: %s, Type: %T\n", i, v, v)
	// }
	// for i := 0; i < len(names); i++ {
	// 	names[i] = "Changes value"
	// 	fmt.Printf("Index: %d", i)
	// }
	// fmt.Print("Names", names)

	// LABEL1:
	// 	for i := 0; i <= 5; i++ {
	// 		for j := 0; j <= 5; j++ {
	// 			if j == 4 {
	// 				continue LABEL1
	// 			}
	// 			fmt.Printf("i is: %d, and j is: %d\n", i, j)
	// 		}
	// 	}

	i := 0
HERE:
	print(i)
	i++
	if i == 5 {
		return
	}
	goto HERE
}
