package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices")

	var fruitList = []string{"Mango", "Apple", "Peach", "Tangrine", "banana"}

	fmt.Printf("Type of FruitList %T", fruitList)

	fruitList = append(fruitList[1:3], "tomato", "pomogrante")
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 0
	highScores[1] = 1
	highScores[2] = 2
	highScores[3] = 3

	highScores = append(highScores, -1)

	fmt.Println(highScores)
	sort.Ints(highScores)
	fmt.Println(highScores)

	//remove based on index
	var index = 2
	highScores = append(highScores[:index], highScores[index+1:]...)
	fmt.Println(highScores)

	for _, value := range highScores {
		fmt.Println(value)
	}

}
