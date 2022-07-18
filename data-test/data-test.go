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

	// array
	var arr1 = [3]int{1, 2, 3}      // length is defined
	arr2 := [...]int{4, 5, 6, 7, 8} // length is inferred

	fmt.Println(arr1)
	fmt.Println(arr2)

	// arrary initialization
	arr3 := [5]int{}              // not initialized
	arr4 := [5]int{1, 2}          // partially initialized
	arr5 := [5]int{1, 2, 3, 4, 5} // fully initialized
	arr6 := [5]int{1: 10, 2: 40}  // initialize only specific elements

	fmt.Println(arr3)
	fmt.Println(arr4)
	fmt.Println(arr5)
	fmt.Println(arr6)

	// lenght of arrary
	fmt.Println(len(arr6))

}
