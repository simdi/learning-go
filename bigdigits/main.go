package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var bigDigits = [][]string{
	{"  000  ", " 0   0 ", "0     0", "0     0", "0     0", " 0   0 ", "  000  "},
	{"  1  ", " 11  ", "  1  ", "  1  ", "  1  ", "  1  ", " 111 "},
	{" 222 ", "2   2", "   2 ", "  2  ", " 2   ", "2    ", "22222"},
	{" 9999", "9   9", "9   9", " 9999", "    9", "    9", "    9"},
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("usage %S <whole number>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	stringsOfDigits := os.Args[1]
	for row := range bigDigits[0] {
		line := ""
		for column := range stringsOfDigits {
			digit := stringsOfDigits[column] - '0'
			if 0 <= digit && digit <= 9 {
				line += bigDigits[digit][row] + " "
			} else {
				log.Fatal("invalid whole number")
			}
		}
		fmt.Println(line)
	}
}
