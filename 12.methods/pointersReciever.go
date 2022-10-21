package main

import (
	"fmt"
	"math"
)

// Declaration of a struct type called Vertex, basically a couple of two float64 numbers
type Vertex struct {
	X, Y float64
}

// Declaration of a method called Abs() that is callable on a Vertex variable and return a float64
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/*
Declaration of a method with a pointer reciever
You can declare methods with pointer receivers.
This means the receiver type has the literal syntax *T for some type T.
For example, the Scale method here is defined on *Vertex.
Methods with pointer receivers can modify the value to which the receiver points (as Scale does here).
Since methods often need to modify their receiver, pointer receivers are more common than value receivers.
With a value receiver, the Scale method operates on a copy of the original Vertex value.
(This is the same behavior as for any other function argument.)
The Scale method must have a pointer receiver to change the Vertex value declared in the main function.
*/
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	fmt.Println("Value of v before of the Scale methods with pointer: ", v)
	v.Scale(10)
	fmt.Println("Value of v after the Scale methods with pointer: ", v)

	fmt.Println(v.Abs())

}

// Without the pointer on Scale the function will modify a local version of v so out the Scale methods v has the same starting value
// In this way Scale modify the value of v in the main
// In the following code we can see the same code but maded by function insted of methods:

/*
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	Scale(&v, 10)
	fmt.Println(Abs(v))
}


functions with a pointer argument must take a pointer
while methods with pointer receivers take either a value or a pointer as the receiver when they are called
The equivalent thing happens in the reverse direction.

There are two reasons to use a pointer receiver.

The first is so that the method can modify the value that its receiver points to.

The second is to avoid copying the value on each method call. This can be more efficient if the receiver is a large struct, for example.
*/
