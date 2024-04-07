package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"` // - means dont send password
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("welcome")
	EncodeJson()
}

func EncodeJson() {
	firstCourse := []course{
		{"React", 299, "scaler", "abc123", []string{"react ", "js"}},
		{"Node", 599, "webD", "abc123", []string{"js ", "node"}},
		{"Mern", 1599, "jsMastery", "def123", nil},
	}

	// tranforming in json
	content, err := json.MarshalIndent(firstCourse, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(content))

}
