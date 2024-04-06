package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://fakestoreapi.com/products"

func main() {
	fmt.Println("welcome")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println("data is :", string(data))
}
