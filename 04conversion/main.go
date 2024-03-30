package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)



func main () {
	fmt.Println("welcome to admin section")
	fmt.Println("please enter the access code ")

	reader := bufio.NewReader(os.Stdin)

	input , _ := reader.ReadString('\n')

	fmt.Println("Thanks for entering  access code" , input)

	// CONVERTED STRING TO INTEGER 
	numCode , err := strconv.ParseFloat(strings.TrimSpace(input) , 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("added 1 to your access code : " ,  numCode + 1)
	}


}