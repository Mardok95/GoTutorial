package main

import (
	"fmt"
)

// You can define your functions or you can go here to find the default functions:
// https://pkg.go.dev/std
// The delcaration of a function is:
// func nameOfFunction(nameOfParameter typeOfParamenter) typeOfResult

func kelvinToCelsius(k float64) float64 { // Function that need a float64 parameter
	k = k - 273.15
	return k
}

func celsiusToFahr(c float64) float64 {
	c = (c*9.0)/5.0 + 32
	return c
}

func kelvinToFahr(k float64) float64 {
	k = k - 273.15
	k = (k*9)/5 + 32
	return k
}
func main() {
	kelvin := 0.0
	celsius := kelvinToCelsius(kelvin)
	fahrenheitCelsius := celsiusToFahr(celsius)
	fahrenheitKelvin := kelvinToFahr(kelvin)
	fmt.Println(kelvin, "° K sono ", celsius, "° C")
	fmt.Println(kelvin, "° K sono ", fahrenheitKelvin, "° far")
	fmt.Println(celsius, "° C sono ", fahrenheitCelsius, "° far")

}
