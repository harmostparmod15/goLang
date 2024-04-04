package main

import "fmt"

func main() {
	fmt.Println("welcome")

	// days := []string{"sun", "mon", "tue", "wed", "thur", "fri"}

	// 1
	// for i := 0; i < len(days); i++ {
	// 	fmt.Println("curr day ", days[i])
	// }

	//2
	// for i := range days {
	// 	fmt.Println("curr day ", days[i])
	// }

	//3
	// for index, value := range days {
	// 	fmt.Printf(" index is %v and values is %v \n", index, value)
	// }

	//4
	val := 5
	for val < 10 {

		if val == 8 {
			break
		}
		fmt.Println("value is ", val)
		val++
	}

}
