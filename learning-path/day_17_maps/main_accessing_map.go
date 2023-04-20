package main

import "fmt"

func main() {

	codes := map[string]string{"en": "English", "fr": "French", "hi": "Hindi"}

	fmt.Println(codes["en"])
	fmt.Println(codes["fr"])
	fmt.Println(codes["hi"])
}
