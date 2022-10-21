/*
An interface type is defined as a set of method signatures.
A value of interface type can hold any value that implements those methods.
Classical type store data (int store decimal number). Interface type say to us
what a type can do
Methods say to us what is the behavior of a type so an interface is declared
with all the methods that a type need to satisfy
*/
package main

import (
	"fmt"
	"strings"
)

// The variable t can contain any kind of value that can meet the methods talk that
// need no argument and return a string
var t interface {
	talk() string
}

type martian struct{}

func (m martian) talk() string {
	return "nack nack"
}

type laser int

func (l laser) talk() string {
	return strings.Repeat("pew", int(1))

}

/*
martian is an empty struct and laser is an int but all of them meet the method
talk so they can be assigned to "t"
*/

func main() {
	t = martian{}
	fmt.Println(t.talk())
	t = laser(3)
	fmt.Println(t.talk())

}

// The multiform variable "t" can take the aspect of martian or laser
// This functionality is called polymorfism
