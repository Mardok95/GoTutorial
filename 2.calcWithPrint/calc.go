package main

import "fmt"

// main è la funzione principale che viene richiamata per prima
func main() {
	fmt.Print("Il mio peso sulla superficie di Marte sarebbe ")
	fmt.Print(60 * 0.3783)
	fmt.Print(" kg ed inoltre avrei ")
	fmt.Print(27 * 365 / 687)
	fmt.Print(" anni.")

	// alternativa sarebbe
	// fmt.Print("Il mio peso sulla superficie di Marte sarebbe")
}
