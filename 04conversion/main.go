package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to out pizza app please rate our products from 1 to 5")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	fmt.Println("Thanks for Rating", input)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1", numRating+1)
	}

}
