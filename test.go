package main

import (
	"fmt"
)

func someFunc() {
	x := 15
	a := &x //memory address
	*a += 17
	fmt.Println(x)
	x = 3
	fmt.Println(*a)
}
