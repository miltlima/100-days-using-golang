package main

import "fmt"

func main() {

	var grades [5]int = [5]int{90, 80, 70, 80, 97}

	for i := 0; i < len(grades); i++ {
		fmt.Println(grades[i])
	}
}
