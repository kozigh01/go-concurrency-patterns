package main

import (
	"fmt"
	"math/rand"
	"time"
)

func example9() {
	rand.Seed(time.Now().UnixNano())
	f := first9(boring9("Joe"), boring9("Ann"), boring9("Dave"), boring9("Karen"))
	fmt.Printf("The first one done: %v\n", f)
	fmt.Println("You're boring; I'm leaving")
}

func boring9(msg string) <-chan string {
	c := make(chan string)
	go func() {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
		c <- fmt.Sprintf("Hello from %s", msg)
	}()
	return c
}

func first9(inputs ...<-chan string) string {
	c := make(chan string)
	f := func(index int) { for { c <- (<- inputs[index]) } }
	for i := range inputs {
		go f(i)
	}
	return <-c
}