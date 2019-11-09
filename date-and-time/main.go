package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	// t := "2019-10-03T06:31:01.998Z"
	fmt.Println("Time", t)
	fmt.Println(t.Format(time.ANSIC))
	myDateString := "2018-01-20"
	formatDate, err := time.Parse("2006-01-02", myDateString)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Parsed time is: %v\n", formatDate)
}
