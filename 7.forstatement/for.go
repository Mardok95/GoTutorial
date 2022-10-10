package main

import (
	"fmt"
	//"time"
)

func main() {
	var count = 10
	for count > 0 {
		fmt.Println(count)
		//time.Sleep(time.Second) serve a rallentare il ciclo di un secondo ad ogni iterazione
		count--
	}
}

// per uscire dal ciclo posso usare anche break
