package main

import (
	"fmt"
	"linechat/httpclient"
	"log"
)

func main() {
	client := httpclient.NewClient()

	// Example GET request
	response, err := client.Get("/")
	if err != nil {
		log.Println("GET Error:", err)
		return
	}
	fmt.Println("GET Response:", string(response))
}
