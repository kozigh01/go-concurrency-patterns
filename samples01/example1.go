package main

import (
	"fmt"
	"math/rand"
	"time"
)

func example1() {
	go boring1("I'm boring...")
	fmt.Println("I'm listening.")
	time.Sleep(2 * time.Second)
	fmt.Println("You're boring: I'm leaving")
}

func boring1(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}