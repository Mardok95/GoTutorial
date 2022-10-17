package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func fillArray(array [10]int) {
	for j := 0; j < len(array); j++ {
		array[j] = rand.Intn(100) + 1
	}
	fmt.Println(array)

}

func main() {
	fmt.Println("Exercise n.16")
	var num [8]int
	var num2 [10]int
	var planets [3]string
	for i := 0; i < len(num); i++ {
		num[i] = rand.Intn(100) + 1
		//fmt.Println(num[i])
	}

	fmt.Println(num)

	planets[0] = "Jupiter"
	planets[1] = "Starun"
	planets[2] = "Mars"

	dwarfs := [5]string{"Pippo", "Pluto", "Paperino", "Tizio", "Caio"}
	fmt.Println(planets)
	fmt.Println(dwarfs)

	sort.StringSlice(planets[:]).Sort()
	fmt.Println(planets)

	fmt.Println("Testing function")
	fillArray(num2)

}
