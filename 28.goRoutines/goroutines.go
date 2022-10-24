package main

import (
	"fmt"
	"time"
)

func main() {
	go sleepyGopher() // Starting point of go routine
	time.Sleep(4 * time.Second)
}

func sleepyGopher() {
	time.Sleep(3 * time.Second)
	fmt.Println("...ronf...")

}

/*
Until now, all the Go code we wrote was as if it were being
performed by a single actor, engaged in his own tasks and oblivious to
what was going on around it. Go programs, however, can
be much more: multiple independent units, each of which is
engaged in its own activity and each of which communicates
with the others to achieve a common goal. In Go, an independently running unit is called a goroutine.
Starting a goroutine does not differ from a function call.
All you have to do is place the reserved word go in front of the
call.

Translated with www.DeepL.com/Translator (free version)

*/
