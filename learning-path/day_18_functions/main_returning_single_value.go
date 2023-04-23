package main

import "fmt"

func addNumbers(a int, b int) string {
	sum := a + b
	return sum
}

func main() {
	sumOfNumbers := addNumbers(2, 3)
	fmt.Print(sumOfNumbers)
}

// the functions return a error , because add numbers are  int and return a string instead.
