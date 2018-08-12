package main

import (
	"fmt"
)

func foo() {
	defer fmt.Println("Done!")
	defer fmt.Println("Are we done?!")
	fmt.Println("Hello world")
}

func tut20() {
	foo()
	fmt.Println("The last")
}
