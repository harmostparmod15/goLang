package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("welcome")

	//1
	var fruits = []string{"apple", "mango", "banana"}
	fmt.Println("fruits are : ", fruits)

	//2
	fruits = append(fruits, "mango", "cherry")
	fmt.Println("more fruits :  ", fruits)

	//3
	fruits = append(fruits[1:3])
	fmt.Println("fruits : ", fruits)

	// make
	hightScore := make([]int, 4)
	hightScore[0] = 87
	hightScore[1] = 34
	hightScore[2] = 67
	hightScore[3] = 98
	hightScore = append(hightScore, 23, 76, 89)

	sort.Ints(hightScore)
	fmt.Println("highscores are : ", hightScore)

}
