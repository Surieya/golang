package main

import "fmt"

func main() {
	fmt.Println("Structs")
	// no inheritance in golang
	surieya := User{"surieya", "s@gmail.com", true, 20}
	fmt.Println(surieya)
	fmt.Printf("%+v\n", surieya)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
