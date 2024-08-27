package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() {
	c := boring("boring")

	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
	}

	joe := boring("Joe")
	mike := boring("Mike")

	for i := 0; i < 5; i++ {
		fmt.Println(<-joe)
		fmt.Println(<-mike)
	}
}

// Generator pattern: function that lunch a goroutine and returns a channel
func boring(msg string) <-chan string { // Returns receive-only channel of strings.
	c := make(chan string)
	go func() { // We launch the goroutine from inside the function.
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c // Return the channel to the caller.
}
