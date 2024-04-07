package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("welcome")
	// PerformGetRequest()
	// PerformPostJsonRequest()
	// PerformPostFormRequest()
}

func PerformGetRequest() {
	const myUrl = "http://localhost:8000/get"
	response, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("status code of request ", response.StatusCode)
	fmt.Println("content length is : ", response.ContentLength)

	data, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	// second method to convert byte data in string
	var responseString strings.Builder
	byteCount, _ := responseString.Write(data)
	fmt.Println(byteCount)
	fmt.Println(responseString.String())
}

func PerformPostJsonRequest() {
	const myUrl = "http://localhost:8000/post"

	// json payload
	requestBody := strings.NewReader(`
{
	"course": "react",
	"price":"3000",
	"platform":"scaler"
}

`)

	// post request to server
	response, err := http.Post(myUrl, "application/json", requestBody)
	if err != nil {
		panic(err)

	}
	defer response.Body.Close()

	// reading response from server
	byteContent, _ := io.ReadAll(response.Body)
	fmt.Println(string(byteContent))
}

func PerformPostFormRequest() {
	const myUrl = "http://localhost:8000/post"

	// form data payload
	data := url.Values{}
	data.Add("firstname", "john")
	data.Add("lastname", "doe")
	data.Add("age", "21")

	// making post request on server
	response, err := http.PostForm(myUrl, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	// reading body  from response
	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}
