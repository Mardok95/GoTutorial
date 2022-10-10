package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Prova di un IF particolare con dichiarazione :=")
	if num := rand.Intn(10); num == 0 { // In questo modo dichiaro il valore direttamente nell'intestazione dell'IF, stessa cosa per lo switch
		fmt.Println("Il numero è pari a 0")
		fmt.Println(num)
	} else if num >= 5 {
		fmt.Println("Il numero è >= 5")
		fmt.Println(num)
	} else {
		fmt.Println("Il numero è <=5 sicuramente e forse anche = 0")
		fmt.Println(num)
	}
	// Se metto qui fmt.Println(num) ottengo errore perché num è stata dichiarata nell'IF

}
