package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("welcome")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1

	fmt.Println("value of dice is : ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("dice value is 1 ")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	case 4:
		fmt.Println("4")
	case 5:
		fmt.Println("5")
	case 6:
		fmt.Println("6 you got one more chance")
	default:
		fmt.Println("you won")
	}

}
