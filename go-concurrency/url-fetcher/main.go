package main

import (
	"fmt"
	"net/http"
	"sync"
)

type Result struct {
	Url        string
	statuscode int
	Err        error
}

func main() {
	urls := []string{
		"https://google.com",
		"https://github.com",
		"https://golang.org",
		"http://youtube1.com",
	}

	results := make(chan Result)

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go fetchurls(url, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		if result.Err != nil {
			fmt.Println("Error :", result.Url, result.Err)
			continue
		}

		fmt.Printf("Fetched the url %s with status %d \n", result.Url, result.statuscode)
	}
}

func fetchurls(url string, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		results <- Result{
			Url: url,
			Err: err,
		}
		return
	}

	defer resp.Body.Close()
	results <- Result{
		Url:        url,
		statuscode: resp.StatusCode,
	}
}
