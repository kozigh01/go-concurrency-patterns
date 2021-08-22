package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Message struct {
	str string
	wait chan bool
}

func example5() {
	rand.Seed(time.Now().UnixNano())
	c := fanin5(boring5("Joe"), boring5("Ann"))
	for i := 0; i < 20; i++ {
		msg1 := <-c
		fmt.Printf("Msg: %q\n", msg1.str)
		msg2 := <-c
		fmt.Printf("Msg: %q\n", msg2.str)
		msg1.wait <-true
		msg2.wait <-true
	}
	fmt.Println("You're boring; I'm leaving")
}

func boring5(msg string) <-chan Message {
	c := make(chan Message)
	w := make(chan bool)
	go func() {
		for i := 0; ; i++ {
			c <- Message { str: fmt.Sprintf("%s %d", msg, i), wait: w }
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			<-w
		}
	}()
	return c
}

func fanin5(input1, input2 <-chan Message) <-chan Message {
	c := make(chan Message)
	go func() { for { c <- (<- input1) } }()
	go func() { for { c <- (<- input2) } }()
	return c
}