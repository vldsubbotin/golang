package main

import (
	"fmt"
	"io/ioutil"
)

type (
	video struct {
		Id          string
		Title       string
		Description string
		Imageurl    string
		Url         string
	}
)

func getVideos() (videos []video) {

	fileBytes, err := ioutil.ReadFile("./videos.json")
	if err != nil {
		panic(err)
	}

	fileContent := string(fileBytes)

	fmt.Printf(fileContent)

	return nil
}
