package main

import "fmt"

func main () {
	fmt.Println("welcome")

	// normal variable
    num := 23;

   // pointer
   var ptr  = &num

   fmt.Println("address of ptr : " , ptr)
   fmt.Println("value of ptr : " , *ptr)


}