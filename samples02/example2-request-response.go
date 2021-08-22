package main

import "fmt"

// key is the idea of a chan that accepts a channel
func example2() {
	reqResp := make(chan chan string)
	quit := make(chan bool)

	go func(quit chan<- bool) {
		response := make(chan string)
		reqResp <- response
		fmt.Printf("recieved response: %q\n", <-response)
		
		quit <- true
	}(quit)

	go func() {
		response := <- reqResp
		fmt.Println("recieved request")
		response <- "howdy"
	}()

	<- quit
	fmt.Println("quiting the program")
}