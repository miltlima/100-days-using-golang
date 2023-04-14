package main

import "fmt"

func main() {
	var x int = 10
	fmt.Println((x < 0) || (x < 200))
	fmt.Println((x < 0) || (x > 200))
}