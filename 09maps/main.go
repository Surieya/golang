package main

import "fmt"

func main() {
	fmt.Println("Maps")
	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["PY"] = "Python"

	fmt.Println(languages)
	fmt.Println(languages["JS"])

	delete(languages, "PY")

	//loops

	for key, value := range languages {
		fmt.Printf("For key %v,value is %v", key, value)
	}

}
