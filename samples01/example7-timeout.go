package main

import (
	"fmt"
	"math/rand"
	"time"
)

func example7() {
	rand.Seed(time.Now().UnixNano())
	
	// timeout for entire conversation
	c := boring7("Jane")
	exit:
	for i := 0; i < 20; i++ {
		select {
		case s := <- c: fmt.Printf("Msg: %q\n", s)
		case <- time.After(600 * time.Millisecond): 
			fmt.Println("You're too slow.")
			break exit
		}		
	}

	// timeout for each iteration
	c = boring7("Joe")
	for i := 0; i < 20; i++ {
		select {
		case s := <- c: fmt.Printf("Msg: %q\n", s)
		case <- time.After(500 * time.Millisecond): fmt.Println("You're too slow.")
		}		
	}

	fmt.Println("You're boring; I'm leaving")
}

func boring7(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
	}()
	return c
}