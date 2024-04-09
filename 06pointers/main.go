package main

import "fmt"

func main() {
	fmt.Println("Pointers")

	// var ptr *int
	// fmt.Println("Value of Pointer is", ptr)

	var myNumber = 23

	// var ptr = &myNumber
	// ptr = &myNumber
	ptr := &myNumber
	fmt.Println("Value of actual Pointer is", ptr)

}
