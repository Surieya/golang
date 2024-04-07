package main

import "fmt"

const LoginToken string = "huahiodhfoiuha" // Public

func main() {
	// fmt.Println("Variables")
	var username string = "Surieya"
	var isLoggedIn bool = true
	var smallVal uint8 = 255
	fmt.Println("Username: ", username)
	fmt.Printf("Variable is of type: %T \n", username)

	fmt.Printf("Variable is of type: %T \n", isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 0.12345
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	var defaultString string
	var defaultInteger int
	fmt.Println(defaultString)
	fmt.Println(defaultInteger)
	fmt.Println(LoginToken)

}
