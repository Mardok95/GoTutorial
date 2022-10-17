package main

// This kind of fucntions are usefull when we need to create a really quick function
import (
	"fmt"
)

var f = func() {
	fmt.Println("anonymous function")
}

func main() {
	fmt.Println("Exercise n.15")
	f()

}

/*
func main(){
	f := func(message string){
		fmt.Println(message)
	}
	f("anonymous function")
}
*/

/*
func main(){
	func() { <- Declare the function
		fmt.Println("anonymous function")
	}() <- Recall the function
}

*/
