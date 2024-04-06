package main

import "fmt"

func main() {
	defer fmt.Println("welcome")
	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three")
	fmt.Println("sir")
	defer fmt.Println("four")

}
