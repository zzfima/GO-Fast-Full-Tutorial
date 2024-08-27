package main

import (
	"fmt"
	"time"
)

func timeOutPerMessage() {
	c := boring("Joe")
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-time.After(1 * time.Second):
			fmt.Println("U r too slow")
			return
		}
	}
}

func timeOutPerAllConversation() {
	tm := time.After(5 * time.Second)
	c := boring("Joe")
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-tm:
			fmt.Println("U r too slow! 5 second passed!")
			return
		}
	}
}
