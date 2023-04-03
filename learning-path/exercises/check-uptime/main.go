package main

import (
    "fmt"
    "net/http"
    "time"
	"os"
	"bufio"
	"sync"
)

func checkURL(url string) (int64, error) {
    startTime := time.Now()
	client := http.Client{
		Timeout: 5 * time.Second,
	}
    resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	return time.Since(startTime).Nanoseconds(), nil
}

func processURLs(urls chan string , wg *sync.WaitGroup) {
	for url := range urls {
		latency, err := checkURL(url)
		if err != nil {
			fmt.Printf("%s is down with error: %v\n", url, err)
			continue
		}
		fmt.Printf("%s is up with latency: d% nanoseconds\n", url, latency)
	}
	wg.Done()
}

func main() {
	file, err := os.Open("urls.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	urls := make(chan string, 10)
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go processURLs(urls, &wg)
	}

	scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		urls <- scanner.Text()
	}
	close(urls)

    wg.Wait()
	if scanner.Err() != nil {
		fmt.Println(scanner.Err())
		os.Exit(1)
	}
}