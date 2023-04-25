package main

import "fmt"

type Student struct {
	name string
	rollNo int
}

func main() {
	st := Student{"Milton", 25}
	fmt.Printf("%+v", st)
}