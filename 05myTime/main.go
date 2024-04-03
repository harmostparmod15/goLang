package main

import (
	"fmt"
	"time"
)

func main () {
	fmt.Println("welcome admin")

	presentTime := time.Now()

	fmt.Println("current time " , presentTime)

	// have to pass this string as it is :-  01-02-2006 15:04:05 Monday
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2020 , time.August , 10, 23, 45,0,0 , time.UTC)

	fmt.Println("date created as : " , createdDate)
}