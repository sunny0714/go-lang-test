package main

import (
	"fmt"
)

func main() {
	mySlice1 := []int{}
	fmt.Println(len(mySlice1))
	fmt.Println(cap(mySlice1))
	fmt.Println(mySlice1)

	mySlice2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println(len(mySlice2))
	fmt.Println(len(mySlice2))
	fmt.Println(mySlice2)

	// create a slice from an arrary
	// The slice can grow to the end of the arrary
	arr1 := [6]int{10, 11, 12, 13, 14, 15}
	mySlice3 := arr1[2:4]
	fmt.Printf("mySlice3 = %v\n", mySlice3)
	fmt.Printf("length = %d\n", len(mySlice3))
	fmt.Printf("capacity = %d\n", cap(mySlice3))

	// create a Slice with the make() function
	mySlice4 := make([]int, 5, 10)
	fmt.Printf("mySlice4 = %v\n", mySlice4)
	fmt.Printf("length = %d\n", len(mySlice4))
	fmt.Printf("length = %d\n", cap(mySlice4))

	mySlice5 := make([]int, 5)
	fmt.Printf("mySlice5 = %v\n", mySlice5)
	fmt.Printf("length = %d\n", len(mySlice5))
	fmt.Printf("capacity = %d\n", cap(mySlice5))

	// appending elements
	myslice6 := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("myslice6 = %v\n", myslice6)
	fmt.Printf("length = %d\n", len(myslice6))
	fmt.Printf("capacity = %d\n", cap(myslice6))

	myslice6 = append(myslice6, 20, 21)
	fmt.Printf("myslice6 = %v\n", myslice6)
	fmt.Printf("length = %d\n", len(myslice6))
	fmt.Printf("capacity = %d\n", cap(myslice6))

	// appending one slice to another slice
	myslice7 := []int{1, 2, 3}
	myslice8 := []int{4, 5, 6}
	myslice9 := append(myslice7, myslice8...)

	fmt.Printf("myslice9=%v\n", myslice9)
	fmt.Printf("length=%d\n", len(myslice9))
	fmt.Printf("capacity=%d\n", cap(myslice9))
}
