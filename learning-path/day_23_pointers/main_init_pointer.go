package main

import "fmt"

func main() {

	s := "hello"
	var b *string = &s
	fmt.Println(b)

	var a = &s
	fmt.Println(a)
	c := &s
	fmt.Println(c)
}