package main

import "fmt"
import "strconv"

func main()  {
	var i int = 90
	var s string = strconv.Itoa(i)

	fmt.Printf("%q", s)
}