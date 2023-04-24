package main

import "fmt"

func modify(s string) {
	s = "world"
}

func main() {

	a := "hello"
	fmt.Println(a)
	modify(a)
	fmt.Println(a)
}