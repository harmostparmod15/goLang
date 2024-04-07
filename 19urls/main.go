package main

import (
	"fmt"
	"net/url"
)

const newUrl = "http://localhost:3000/learn?course=react&page=1"

func main() {
	fmt.Println("welcome")

	result, err := url.Parse(newUrl)
	if err != nil {
		panic(err)
	}

	// printing diff parts of url
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	// extracting query params from url
	qParams := result.Query()
	fmt.Println("query", qParams)
	course := qParams.Get("course")
	fmt.Println("course query ", course)

	// making url string from parts  have to pass url as referrence its imp.
	partOfUrl := &url.URL{
		Scheme:   "http",
		Host:     "localhost:3000",
		Path:     "/learn",
		RawQuery: "page=1",
	}

	anotherUrl := partOfUrl.String()
	fmt.Println("url : ", anotherUrl)

}
