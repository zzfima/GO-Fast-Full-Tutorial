package main

import (
	"fmt"
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

// Generator: function that lunch a goroutine and returns a channel
func boring(str string) <-chan string { //returns receive-only channel of strings
	c := make(chan string)

	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second)
			c <- str + " " + fmt.Sprintf("%d", i)
		}
	}()

	return c
}
