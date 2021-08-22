package main

import (
	"fmt"
	"math/rand"
	"time"
)

func example6() {
	rand.Seed(time.Now().UnixNano())
	c := fanin6(boring6("Joe"), boring6("Ann"))
	for i := 0; i < 20; i++ {
		fmt.Printf("Msg: %q\n", <-c)
	}
	fmt.Println("You're boring; I'm leaving")
}

func boring6(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
	}()
	return c
}

func fanin6(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <- input1: c <- s
			case s := <- input2: c <- s
			}
		}
	}()
	return c
}