package main

import "fmt"

func main() {
	fmt.Println("welcome")
	greet()
	var result int = sum(3, 5)
	fmt.Println("the sum of 2 numbers is : ", result)

	var proResult int = adder(4, 5, 4, 2, 5, 7, 64, 3)
	fmt.Println("pro result is : ", proResult)
}

func greet() {
	fmt.Println("hello and welcome to go lang")
}

func sum(a int, b int) int {
	return a + b
}

// pro function
func adder(values ...int) int {
	total := 0

	for _, val := range values {
		total += val
	}
	return total
}
