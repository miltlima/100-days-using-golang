package main

import "fmt"
import "reflect"

func main() {

	fmt.Printf("Type: %v \n", reflect.TypeOf(1000))
	fmt.Printf("Type: %v \n", reflect.TypeOf("Milton"))
	fmt.Printf("Type: %v \n", reflect.TypeOf(12.4))
	fmt.Printf("Type: %v \n", reflect.TypeOf(true))
}