package main

import "fmt"

func main() {

	var fruits [2]string = [2]string{"apples", "oranges"}
	fmt.Println(fruits)

	marks := [3]int{10, 20, 30}
	fmt.Println(marks)

	names := [...]string{"Rachel", "Phoebe", "Monica"}
	fmt.Println(names)
}
