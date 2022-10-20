// You can read the assignment here: https://go.dev/tour/flowcontrol/8

package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	var z = 1.0
	var z1 float64
	var z2 float64
	for i := 0; i < 50; i++ {
		z1 = z
		z -= (z*z - x) / (2 * z)
		z2 = z - z1
		/*
			fmt.Println("This is z: ",z)
			fmt.Println("This is z1: ",z1)
			fmt.Println("This is z2: ",z2)
			fmt.Println("This is i: ",i)
		*/
		if z2 >= 0 && z2 <= 0.01 {
			fmt.Println("This is the last execution: ", i)
			i = 100
		}
	}
	return z

}

func main() {
	x := 257.0
	fmt.Println("The sqrt of ", x, " is: ", Sqrt(x))

}
