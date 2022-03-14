package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func readData() {
	content, err := ioutil.ReadFile("./hello")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("File contents: %s", content)

}
