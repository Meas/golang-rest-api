package main

import (
	"fmt"
	"sync"
	"time"
)

var wgg sync.WaitGroup

func tut18() {
	wg.Add(1)
	go say("hey")
	wg.Add(1)
	go say("there")

	wg.Add(1)
	go say("Hi")
	wg.Wait()
}

func cleanUp() {
	defer wg.Done()
	if r := recover(); r != nil {
		fmt.Println("Recovered in cleanup", r)
	}
}

func say(s string) {
	defer cleanUp()
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
		if i == 2 {
			panic("Oh dear, a 2")
		}
	}
}
