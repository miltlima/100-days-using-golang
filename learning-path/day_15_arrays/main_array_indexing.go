package main

import "fmt"

func main() {

	var grades [5]int = [5]int{90, 80, 70, 80, 97}
	fmt.Println(grades)
	grades[1] = 100
	fmt.Println(grades)
}
