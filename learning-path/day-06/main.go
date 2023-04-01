package main

import "fmt"

func main() {
	const name = "Milton Lima"
	const is_muggle = false
	const age = 39

	fmt.Printf("%v: %T \n", name, name)
	fmt.Printf("%v: %T \n", is_muggle, is_muggle)
	fmt.Printf("%v: %T \n", age, age)
}