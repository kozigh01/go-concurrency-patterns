package main

import (
	"fmt"
	"math/rand"
	"time"
)

func example8() {
	rand.Seed(time.Now().UnixNano())

	quit := make(chan string)
	t := time.After(3 * time.Second)
	c := boring8("Joe", quit)

	exit:
	for i := 0; i < 20; i++ {
		select {
		case s := <- c: fmt.Printf("Msg: %q\n", s)
		case <- t: 
			quit <- "Bye!"
			fmt.Printf("Joe says: %q\n", <-quit)
			break exit
		}		
	}

	fmt.Println("You're boring; I'm leaving")
}

func boring8(msg string, quit chan string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			select {
			case <- quit:
				fmt.Println("Quitting.")
				// cleanup()
				quit <- "See ya!"
				return
			case c <- fmt.Sprintf("%s %d", msg, i):
				time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			}
		}
	}()
	return c
}