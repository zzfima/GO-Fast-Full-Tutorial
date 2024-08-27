package main

import (
	"fmt"
	"math/rand"
	"time"
)

func quit() {
	quit := make(chan bool)
	limit := rand.Intn(10)
	fmt.Println(limit)

	c := boringQuit("Joe", quit)
	for i := limit; i >= 0; i-- {
		fmt.Println(<-c)
	}
	quit <- true
}

func boringQuit(msg string, quit <-chan bool) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			select {
			case c <- fmt.Sprintf("%s: %d", msg, i):
				// do nothing
			case <-quit:
				return
			}
		}
	}()
	return c
}
