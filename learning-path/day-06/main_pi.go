package main

import "fmt"

const PI float64 = 3.14 

func main() {
	var radius float64 = 6.0
	var area float64
	area = PI * radius * radius
	fmt.Println("Area of Circle is : ", area)
}