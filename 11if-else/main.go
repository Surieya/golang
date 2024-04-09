package main

import "fmt"

func main() {
	fmt.Println("If Else")

	loginCount := 23

	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else {
		result = "Nothing"
	}
	fmt.Println(result)

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is not less than 10")
	}

}
