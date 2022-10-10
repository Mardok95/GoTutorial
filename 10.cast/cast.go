package main

import (
	"fmt"
)

func main() {
	fmt.Println("Example of cast")
	var anno = 1995
	var variabile = 10.05
	// If i try some operations between different types i'll get an error
	var result = anno / int(variabile)
	fmt.Println("The result is: ", result)

	// pay attention to the cast because if you use int(num) you'll lost the decimal number

	var var2 byte = 33
	fmt.Println("This is the cast between number and string")
	fmt.Println(string(var2))

	// You can cast into string with string(). Remember that the types in Go are static

}
