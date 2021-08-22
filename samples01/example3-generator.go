package main

import (
	"fmt"
	"math/rand"
	"time"
)


func example3() {
	joe := boring3("Joe")
	ann := boring3("Ann")
	for i := 0; i < 5; i++ {
		fmt.Printf("Msg: %q\n", <-joe)
		fmt.Printf("Msg: %q\n", <-ann)
	}
	fmt.Println("You're boring; I'm leaving")
}

func boring3(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
	}()
	return c
}