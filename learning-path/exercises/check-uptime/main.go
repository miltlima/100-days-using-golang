package main

import (
    "fmt"
    "net/http"
    "time"
	"os"
	"bufio"
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

func main() {
	file, err := os.Open("urls.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	var wg sync.WaitGroup

	scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			latency, err := checkURL(url)
			if err != nil {
				fmt.Printf("%s is down with error: %v\n", url, err)
				return
        	}
        	fmt.Printf("%s is up with latency: %d nanoseconds\n", url, latency)
    	}(scanner.Text())
	}
	wg.Wait()
    if scanner.Err() != nil {
        fmt.Println(scanner.Err())
        os.Exit(1)
    }
}