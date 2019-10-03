package main

import (
	"fmt"
	"time"
)

func main() {
	myDateString := "2018-01-20"
	formatDate, err := time.Parse("2006-01-02", myDateString)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Parsed time is: %v\n", formatDate)
}
