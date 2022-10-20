package main

import "fmt"

// || significa OR, && signfica AND

func main() {
	var command = "vai est"
	if command == "vai est" {
		fmt.Println("Stai andando ad est")
	} else if command == "vai dentro" {
		fmt.Println("Non stai andando ad est")
	} else {
		fmt.Println("Errore")
	}

}

/*
func main() {
	fmt.Println("C'è l'ingresso di una caverna e un sentiero che va a est.")
	var command = "vai dentro"
	switch command {
	case "vai est":
		fmt.Println("Ti dirigi verso la montagna.")
	case "entra caverna", "vai dentro":
		fmt.Println("Ti trovi in una caverna buia.")
	case "leggi cartello":
		fmt.Println("Sul cartello c'è scritto 'Vietato ai minori'.")
	default:
		fmt.Println("Non ho capito.")
	}

}

You can use switch with no statement and the meaning is switch true, so it's basically an if-then-else condition
package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}






*/
