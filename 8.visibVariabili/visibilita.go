package main

import (
	"fmt"
	"math/rand"
)

// In Go la visibilità, in genere, si esaurisce all’interno di una coppia diparentesi graffe, { }.
// Se dichiaro qui una variabile, viene dichiarata per tutto il package, quindi se ci sono più funzioni
// allora tutte le funzioni vedranno quella variabile
func main() {
	fmt.Println("Parliamo di visibilità delle variabili")
	var count = 0
	for count < 10 { // nuovo campo di visibilità
		var num = rand.Intn(10) + 1
		fmt.Println(num)
		count++
	} // termina il campo di visibilità aperto con il for

	// count è stata dichiarata nel main quindi tutto ciò che è nel main vede count
	// num è stata dichiarata nel for, quindi al di fuori del for la variabile num "non esiste"

}
