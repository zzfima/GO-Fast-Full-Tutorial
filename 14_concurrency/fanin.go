package main

import "fmt"

func fanInClient() {
	c := fanInWithSwitch(boring("Joe"), boring("Mike"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
}

//Multiplexer(fan in) pattern. Combine channels into one channel
func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()

	go func() {
		for {
			c <- <-input2
		}
	}()

	return c
}

//Multiplexer(fan in) pattern. Combine channels into one channel
func fanInWithSwitch(input1, input2 <-chan string) <-chan string {
	c := make(chan string)

	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()

	return c
}
