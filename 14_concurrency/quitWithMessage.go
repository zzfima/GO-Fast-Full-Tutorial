package main

import (
	"fmt"
	"math/rand"
	"time"
)

func cleanup() {
	fmt.Println("cleaning here...")
}

func quitWithMessage() {
	quit := make(chan string)
	limit := rand.Intn(10)
	fmt.Println(limit)
	c := boringQuitStr("Joe", quit)

	for i := limit; i >= 0; i-- {
		fmt.Println(<-c)
	}

	quit <- "Bye!"
	fmt.Println("incoming server message ", <-quit)
}

func boringQuitStr(msg string, quit chan string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			select {
			case c <- fmt.Sprintf("%s: %d", msg, i):
				// do nothing
			case incomingMsg := <-quit:
				fmt.Println("incoming client message: ", incomingMsg)
				cleanup()
				quit <- "See you!"
				return
			}
		}
	}()
	return c
}
