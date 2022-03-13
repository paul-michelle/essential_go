package main

import (
	"fmt"
	"net/http"
)

func getContentType(url string, outputChannel chan string) {
	resp, err := http.Get(url)
	if err != nil {
		outputChannel <- fmt.Sprintf("error occurred: %s\n", err)
		return
	}
	defer resp.Body.Close()

	contentType := resp.Header.Get("content-type")
	outputChannel <- fmt.Sprintf("Content-type of response from %s is %s\n", url, contentType)
}

func main() {
	urls := []string{
		"https://golang.org",
		"https://api.github.com",
		"https://httpbin.org/xml",
	}

	channel := make(chan string)
	for _, url := range urls {
		go getContentType(url, channel)
	}

	for range urls {
		fmt.Println(<-channel)
	}
}
