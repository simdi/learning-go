package main

import (
	"bytes"
	"fmt"
)

type path []byte

func (p *path) TruncateAtFinalSlash() {
	fmt.Println("byte ", bytes.LastIndex(*p, []byte("/")))
	i := bytes.LastIndex(*p, []byte("/"))
	if i >= 0 {
		*p = (*p)[0:i]
	}
}

func main() {
	pathName := path("/usr/bin/tso") // Conversion from string to path.
	fmt.Printf("Path type %T\n", pathName)
	pathName.TruncateAtFinalSlash()
	fmt.Printf("%s\n", pathName)
}
