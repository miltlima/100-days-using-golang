package main

import "fmt"

func main() {

	var a, b int = 10, 20
	switch {
		case a + b == 30:
			fmt.Println("equal to 30")
		case a + b <= 30:
			fmt.Println("less than or equal to 30")
		default:
			fmt.Println("greater than 30")
	}
}