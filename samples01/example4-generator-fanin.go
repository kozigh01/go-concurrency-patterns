package main

import (
	"fmt"
	"math/rand"
	"time"
)

func example4() {
	rand.Seed(time.Now().UnixNano())
	c := fanin4(boring4("Joe"), boring4("Ann"))
	for i := 0; i < 20; i++ {
		fmt.Printf("Msg: %q\n", <-c)
	}
	fmt.Println("You're boring; I'm leaving")
}

func boring4(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
	}()
	return c
}

func fanin4(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() { for { c <- (<- input1) } }()
	go func() { for { c <- (<- input2) } }()
	return c
}