package main

import (
	"fmt"
	"time"
)

type Ball1 struct { hits int }

func example1() {
	table := make(chan *Ball1)
	go player("ping", table)
	go player("pong", table)

	table <- new(Ball1)
	time.Sleep(1 * time.Second)
	<- table

	// can use panic to see the stack traces
	// panic("show me the stacks")
}

func player(name string, table chan *Ball1) {
	for {
		ball := <-table
		ball.hits++
		fmt.Println(name, ball.hits)
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}