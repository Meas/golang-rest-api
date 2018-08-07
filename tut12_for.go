package main

import (
	"fmt"
)

func tut12() {
	x := 5
	for {
		fmt.Println(x)
		x += 3
		if x > 25 {
			break
		}
	}
}
