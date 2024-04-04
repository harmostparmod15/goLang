package main

import "fmt"

func main() {
	fmt.Println("welcome")

	languages := make(map[string]string)
	languages["JS"] = "javascript"
	languages["RB"] = "ruby"
	languages["PY"] = "python"

	fmt.Println("list of all languages", languages)
	fmt.Println("js shorts for ", languages["JS"])

	// loops
	for key, value := range languages {
		fmt.Printf("for key %v , value is %v \n: ", key, value)
	}

}
