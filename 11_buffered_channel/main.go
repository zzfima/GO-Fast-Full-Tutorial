package main

import (
	"fmt"
	"time"
)

func main() {
	var chBuff = make(chan int, 5)
	go generate(chBuff)
	fmt.Println(<-chBuff)
	fmt.Println(<-chBuff)
	fmt.Println(<-chBuff)
	fmt.Println(<-chBuff)
	fmt.Println(<-chBuff)
}

func generate(chBuff chan int) {
	time.Sleep(time.Second)
	chBuff <- 2
	time.Sleep(time.Second)
	chBuff <- 3
	time.Sleep(time.Second)
	chBuff <- 4
	time.Sleep(time.Second)
	chBuff <- 5
	time.Sleep(time.Second)
	chBuff <- 6
	time.Sleep(time.Second)
}
