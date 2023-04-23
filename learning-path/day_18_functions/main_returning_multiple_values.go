package main

import "fmt"

func operation(a int, b int) (sum int, diff int) {
	sum = a + b
	diff = a - b
	return sum, diff
}

func main() {
	sum, difference := operation(20, 10)
	fmt.Println(sum, " ", difference)
}
