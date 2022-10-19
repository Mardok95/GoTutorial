package main

import (
	"fmt"
)

func main() {
	fmt.Println("Exercise n.18 about structures")
	var curiosity struct {
		lat  float64
		long float64
	}
	curiosity.lat = -4.5895
	curiosity.long = 137.4417
	fmt.Println("Latitudine: ", curiosity.lat, " Longitudine: ", curiosity.long)
	fmt.Println(curiosity)

}

/*
If you need to reuse a lot of time a kind of struct you can declare a struct type
type location struct {
	lat float64
	long float64
}
var position1 location
position1.lat = -14.5684
position1.long = 354.4734
var position2 location
position2.lat = -40.1452
position2.long = 250.1147
*/
