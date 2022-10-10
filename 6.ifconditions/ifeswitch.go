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





*/
