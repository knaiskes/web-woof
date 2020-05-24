package main

import (
	"net/http"
	"sync"
	"fmt"
)

var urls = []string {
}

func getStatus(url string, wg *sync.WaitGroup) (string, error) {
	req, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	wg.Done()
	fmt.Printf("url: %s Status Code: %s\n", url, req.Status)
	return req.Status, nil
}

func printStatus() {
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go getStatus(url, &wg)
	}
	wg.Wait()
}

func main() {
	printStatus()
}
