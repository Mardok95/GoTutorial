package main

import "fmt"

func main() {
	fmt.Println("New kind of for")
	sum := 0
	// Classic for
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	// The init and the post are optional
	for sum <= 100 {
		sum += sum
	}
	fmt.Println(sum)
	// The FOR statement in Go can be used as WHILE statement of C/Java ect..
}
