package main

import (
	"fmt"
	"time"
)

func main() {
	//lockSituation()
	noLockSituation()
}

func noLockSituation() {
	//channel is thread safe, hold data and listen for data
	var ch chan int = make(chan int)
	go addToChannel(ch)
	fmt.Println(<-ch)
}

func addToChannel(ch chan int) {
	time.Sleep(2 * time.Second)
	ch <- 1
}

func lockSituation() {
	//channel is thread safe, hold data and listen for data
	var ch chan int = make(chan int)

	//add val 1 to channel. and wait until somebody will read from it. Forever.
	//So it lock current thread. Deadlock
	ch <- 1

	//Can not reach here
	//retrieve value from channel
	var i = <-ch

	fmt.Println(i)
}
