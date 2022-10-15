package main

import "fmt"

type celsius float64 // dichiaro un tipo celsius che è un float64
type kelvin float64

func kelvinToCelsius(k kelvin) celsius { // funzione classica che usa il tipo kelvin ed il tipo celsius
	return celsius(k - 273.15)
}

func (k kelvin) celsius() celsius { // Il metodo celsius di tipo kelvin
	return celsius(k - 273.15)
}

func main() {
	var k kelvin = 294.0
	var c celsius
	c = kelvinToCelsius(k) // richiamo la funzione kelvinToCelsius
	fmt.Println("Ecco celsius come funzione C: ", c)
	c = k.celsius() // richiamo il metodo celsius
	fmt.Println("Ecco celsius come metodo C: ", c)

}
