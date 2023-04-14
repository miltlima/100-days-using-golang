package main

import "fmt" 

func main() {

	fruit := "grapes"
	if fruit == "apple" {
		fmt.Println("I love apples")
	} else if fruit == "orange" {
		fmt.Println("Oranges are not apples")
	} else {
		fmt.Println("no appetite")
	}
}