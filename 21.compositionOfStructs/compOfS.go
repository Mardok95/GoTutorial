package main

import "fmt"

/*
For a complex struct you can use one struct with all the details inside
but it can be complex in the future, like this one:
type report strct {
	sol int
	high, low float64
	lat, long float64
}
The code is not speaking
*/

type report struct {
	sol         int
	temperature temperature
	location    location
}
type temperature struct {
	high, low celsius
}
type location struct {
	lat, long float64
}
type celsius float64

func main() {
	bradbury := location{-4.5895, 137.4417}
	t := temperature{high: -1.0, low: -78.0}
	report := report{sol: 15, temperature: t, location: bradbury}
	fmt.Printf("%+v\n", report)
	fmt.Printf("un tiepido %v° C\n", report.temperature.high)

}
