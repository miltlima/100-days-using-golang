package main

import "fmt"

func main() {

	arr := [5]int{10, 20, 30, 40, 50}
	i := 2
	fmt.Println(arr)

	slice_1 := arr[:i]
	slice_2 := arr[i+1:]

	new_slice := append(slice_1, slice_2...)

	fmt.Println(new_slice)
}
