package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {

	fmt.Println("welcome")
	parmod := User{"parmod", "parmod@gmail.com", true, 12}

	fmt.Println("user created ", parmod)
	//  more details
	fmt.Printf("user created %+v \n", parmod)

}
