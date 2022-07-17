package main

import (
	"fmt"
)

// Declaring a constant
const PI = 3.14
const A int = 1

const (
	C int = 3
	D     = 5
	E     = "Hello!"
)

func main() {
	const B int = 2
	fmt.Println(PI)
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
	fmt.Println(E)

	var i string = "Hello"
	var j int = 15

	fmt.Printf("i has value: %v and type: %T\n", i, i)
	fmt.Printf("j has value: %v and type: %T\n", j, j)
}
