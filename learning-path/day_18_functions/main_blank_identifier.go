package main

import "fmt"

func f() (int, int) {
	return 42, 53
}

func main() {
	v, _ := f()
	fmt.Println(v)
}
