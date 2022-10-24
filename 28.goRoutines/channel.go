package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Exercise n.28 on channel in goroutines")
	c := make(chan int) // This channel can send and recieve int value with the arrow operator <-
	for i := 0; i < 5; i++ {
		go sleepyGopher(i, c)
	}
	for i := 0; i < 5; i++ {
		gopherId := <-c
		fmt.Println("The goroutin ", gopherId, " stopped to sleep")
	}

}

func sleepyGopher(id int, c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("....", id, " ronf ....")
	c <- id

}

/*
A channel can be used to send value in safety from a goroutine to another.
Channel are a type and can be used as a variable in go.
You can create a channel with the statement make
*/
