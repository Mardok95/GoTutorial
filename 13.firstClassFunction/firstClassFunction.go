package main

import (
	"fmt"
	"math/rand"
)

type kelvin float64

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

// First class means that function can be insert in every space of code where you can put int or string

func main() {
	fmt.Println("Ex n. 13")
	fakeS := fakeSensor()
	fmt.Println("This is the assignment of a function to a varabile: ", fakeS)

}
