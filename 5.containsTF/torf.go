package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Sei in una caverna buia")
	var command = "vai fuori"
	var exit = strings.Contains(command, "fuori")
	fmt.Println("Esci dalla caverna, true o false? --> ", exit)

}
