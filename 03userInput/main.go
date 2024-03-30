package main

import (
	"bufio"
	"fmt"
	"os"
)



func main() {
	var welcome string = "welcome to user input section"
	 fmt.Println(welcome)

    // TAKING INPUT FROM USER
	 reader := bufio.NewReader(os.Stdin)
	 fmt.Println("Enter your secret code  : ")

	 // COMMA OK SYNTAX
	 input , err := reader.ReadString('\n')
	 fmt.Println( " your entered code is :" , input)
	 fmt.Println("error occured" , err)

}

