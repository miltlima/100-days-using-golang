package main

import "fmt"

type Student struct {
	name string
	rollNo int
	marks []int
	grades map[string]int
}

func main () {
	st := new(Student)
	fmt.Printf("%+v", st)
}