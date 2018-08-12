package main

import (
	"fmt"
	"sync"
)

var waitGroup sync.WaitGroup

func fooFunc(c chan int, someValue int) {
	defer waitGroup.Done()
	c <- someValue * 5
}

func tut22() {
	fooVal := make(chan int, 10)

	for i := 0; i < 10; i++ {
		waitGroup.Add(1)
		go fooFunc(fooVal, i)
	}
	waitGroup.Wait()
	close(fooVal)

	for item := range fooVal {
		fmt.Println(item)
	}

	/* go fooFunc(fooVal, 5)
	go fooFunc(fooVal, 3)

	v1, v2 := <-fooVal, <-fooVal

	fmt.Println(v1, v2) */
}
