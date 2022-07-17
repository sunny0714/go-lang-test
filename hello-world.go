// This is a single line comment
package main

import "fmt"

func main() {
	/* This is a
	   multi line comment */
	fmt.Println("Hello, World!")
	// fmt.Println("Comment to prevent conde execution")

	var student1 string = "John" // type is string
	var student2 = "Jane"        // type is inferred
	x := 2                       // type is inferred

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)
}
