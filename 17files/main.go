package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	fmt.Println("welcome")

	// content := "this will be content of the file"

	// creating file
	// file, err := os.Create("./test.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// writing to file
	// len, err := io.WriteString(file, content)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("file created of length :", len)
	// defer file.Close()
	readFile("./test.txt")

}

// function to read from file
func readFile(fileName string) {
	dataByte, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Println("content :", string(dataByte))

}
