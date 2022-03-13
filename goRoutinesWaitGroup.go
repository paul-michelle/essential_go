package main

import (
	"fmt"
	"net/http"
	"sync"
)

func getContentType(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("error occurred: %s\n", err)
	}
	defer resp.Body.Close()

	return resp.Header.Get("content-type")
}

func main() {
	urls := []string{
		"https://golang.org",
		"https://api.github.com",
		"https://httpbin.org/xml",
	}
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			contentType := getContentType(url)
			fmt.Println(contentType)
			wg.Done()
		}(url)
	}
	wg.Wait()
}