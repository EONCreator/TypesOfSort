package main

import (
	"fmt"
)

func main() {
	x := 4
	p := &x

	x = 8

	fmt.Println(x)
	fmt.Println(*p)
}
