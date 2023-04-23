package main

import "fmt"

func printDetails(student string, subjects ...string) {
	fmt.Println("hey ", student, ", here are your subjects - ")
	for _, sub := range subjects {
		fmt.Print("%s, ", sub)
	}
}

func main() {
	printDetails("Milton", "Technology", "Biology")
}
