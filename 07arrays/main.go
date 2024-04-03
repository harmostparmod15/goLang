package main

import "fmt"

func main() {
	fmt.Println("welcome")

	// declaring array
	var numbers [4]int

	// initializing array
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40

	fmt.Println("numbers are : ", numbers)

	// declaring and initializing array
	var fruits = [3]string{"apple", "mango", "pear"}

	fmt.Println("fruits are : ", fruits)
}
