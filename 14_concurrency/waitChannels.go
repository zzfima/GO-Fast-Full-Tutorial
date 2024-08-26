package main

import (
	"fmt"
	"time"
)

type Message struct {
	str  string
	wait chan bool
}

func waitChannels() {

	waitForIt := make(chan bool)
	c := make(chan Message)

	go func() {
		for i := 0; i < 5; i++ {
			msg1 := <-c
			fmt.Println(msg1.str)
			msg2 := <-c
			fmt.Println(msg2.str)

			msg1.wait <- true
			msg1.wait <- true
		}
	}()

	c <- Message{str: "hello", wait: waitForIt}
	time.Sleep(time.Second)
	c <- Message{str: "good bye", wait: waitForIt}
	<-waitForIt
}
