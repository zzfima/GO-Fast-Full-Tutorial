package main

import "fmt"

//receive num from right, increment, send to left
func receiveIncrementSend(left chan<- int, right <-chan int) {
	left <- 1 + <-right
}

func daisyChain() {
	const n = 10000
	leftmost := make(chan int)
	right := leftmost
	left := leftmost

	for i := 0; i < n; i++ {
		right = make(chan int)
		go receiveIncrementSend(left, right)
		left = right
	}

	go func(c chan int) {
		c <- 1
	}(right)

	fmt.Println(<-leftmost)
}
