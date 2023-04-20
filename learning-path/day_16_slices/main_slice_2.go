package main

import "fmt"

func main() {

	arr := [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	slice := arr[1:8]
	fmt.Println(slice)

	sub_slice := slice[0:3]
	fmt.Println(sub_slice)
}
