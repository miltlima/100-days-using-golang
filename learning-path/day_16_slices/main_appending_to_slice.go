package main

import "fmt"

func main() {

	arr := [5]int{10, 20, 30, 40, 50}

	slice := arr[:2]

	arr_2 := [5]int{5, 15, 25, 35, 45}

	slice_2 := arr_2[:2]

	new_slice := append(slice, slice_2...)

	fmt.Println(new_slice)

}
