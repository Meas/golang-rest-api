package main

/* var wg sync.WaitGroup */

/* func tut18() {
	wg.Add(1)
	go say("hey")
	wg.Add(1)
	go say("there")


	wg.Add(1)
	go say("Hi")
	wg.Wait()
}

func say(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done()
} */
