package main

import (
	"fmt"
	"math/rand"
	"time"
)

func example3() {
	rand.Seed(time.Now().UnixNano())
	a, b := make(chan string), make(chan string)

	coinFlip := rand.Intn(2) == 0
	if coinFlip {
		a = nil
		fmt.Println("nil a")
	} else {
		b = nil
		fmt.Println("nil b")
	}

	go func() { a <- "a"}()
	go func() { b <- "b"}()

	// nil channel blocks indefinitely, so won't be selected
	select {
	case s := <-a:
		fmt.Println("got", s)
	case s := <-b:
		fmt.Println("got", s)
	}
}