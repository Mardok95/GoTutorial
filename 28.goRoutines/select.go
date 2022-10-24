package main

import (
	"fmt"
)

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	fmt.Println("Exercise about select in goroutines")
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)

}

/*
To manage differt goroutines that produce different value you can use the select statement.
It's similiar to the switch statement. Every case in the select contains a send and a recieve on the channel.
Select wait until the case is ready to start and then execute the code in the case.
It is as if select observes both channels at the same time and takes an action at the moment when there is
something occurs on one of them.
The default statement is run if no other case is ready.
*/
