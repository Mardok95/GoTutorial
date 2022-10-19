package main

import (
	"fmt"
)

func main() {
	fmt.Println("Esercise n.17")
	temperature := map[string]int{
		"Terra": 15, // Couples of string key and int value
		"Marte": -65,
	}
	temp := temperature["Terra"]
	fmt.Printf("The temperature on the Earth is %v ° C.\n", temp)
	temperature["Venere"] = 464
	fmt.Println("All the temperature are the following:")
	fmt.Println(temperature)
}
