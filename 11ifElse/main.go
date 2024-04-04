package main

import "fmt"

func main() {
	fmt.Println("welcome ")

	loginCount := 3
	var result string

	if loginCount < 2 {
		result = "dead user"
	} else if loginCount == 3 {
		result = "miracle user"
	} else {
		result = "regular user"
	}

	fmt.Println(result)

	// declaring and checking variable on the go
	if num := 6; num < 5 {
		fmt.Println("less than 5")
	} else {
		fmt.Println("greater than 5")
	}

}
