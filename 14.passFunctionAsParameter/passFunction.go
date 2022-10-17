package main

import (
	"fmt"
	"math/rand"
	"time"
)

type kelvin float64

func measureTemp(samples int, sensor func() kelvin) { // This function accept another function as second parameter
	for i := 0; i < samples; i++ {
		k := sensor()
		fmt.Printf("%v° K\n", k)
		time.Sleep(time.Second)
	}
}

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)

}

func main() {
	measureTemp(3, fakeSensor) // To call the function i need to pass, as second parameter, another function
}

/*

 We can declare a type of function to make code more reliable and readable:
 type sensor func() kelvin

 in this way we can type this delcaration:
 func measureTemp(samples int, s sensor)

*/
