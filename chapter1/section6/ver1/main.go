package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	url, err := url.Parse("https://localhost:8080/?page=1&page=2&a=1")
	if err != nil {
		log.Fatal(err)
	}
	queryMap := url.Query()

	fmt.Println(queryMap["page"])
	fmt.Println(queryMap["a"])
	fmt.Println(queryMap["b"])
}
