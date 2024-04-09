package main

import "fmt"

func main() {
	fmt.Println("Arrays")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "PineApple"
	fruitList[3] = "Pinees"

	fmt.Println(fruitList)
}
