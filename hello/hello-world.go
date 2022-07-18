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

	var a string
	var b int
	var c bool

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// Value assignment after declaration
	var student3 string
	student3 = "Jane"
	fmt.Println(student3)

	// Difference between var and :=
	// var can be used inside and outside of functions, and variable declaration and value assignment can be done separately
	// := can only be used inside function, and variable declaration and value assignment must be done in the same line

	// mutiple variable declaration
	// If use "type" keyword, it's only possible to declare one type of variable per line

	var d, e int = 1, 3
	var f, g = 6, "Hello"
	h, i := 7, "World!"

	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)

	// varaible declaration in a block

	var (
		j int
		k int    = 1
		l string = "success"
	)

	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)
}
