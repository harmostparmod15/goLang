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
	parmod.GetStatus()
	parmod.NewMail()

	// fmt.Println("user created ", parmod)
	//  more details
	// fmt.Printf("user created %+v \n", parmod)

}

func (u User) GetStatus() {
	fmt.Println("user active :", u.Status)
}

// its a copy , will not change original struct
func (u User) NewMail() {
	u.Email = "test@gmail.com"
	fmt.Println("new email is ", u.Email)
}
