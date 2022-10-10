package main

import "fmt"

func main() {
	const lightSpeed = 299792 // dichiaro una costante che non può essere modificata
	var distance = 56000000   // dichiaro una variabile e gli assegno un valore
	fmt.Println(distance/lightSpeed, "secondi")
	distance = 401000000 // cambio il valore assegnato
	fmt.Println(distance/lightSpeed, "secondi")

	// Posso dichiarare le variabili anche in questo modo
	var (
		variabile1 = 100
		variabile2 = 300
		variabile3 = 400
	)

	// Possiamo anche dichiararle così
	// var variabile1, variabile2, variabile3 = 100, 300, 400

	fmt.Println("Il valore della prima var e' ", variabile1)
	fmt.Println("Il valore della seconda var e' ", variabile2)
	fmt.Println("Il valore della terza var e' ", variabile3)
}
