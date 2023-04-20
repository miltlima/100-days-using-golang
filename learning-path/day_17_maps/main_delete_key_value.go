package main

import "fmt"

func main() {

	codes := map[string]string{"en": "English", "fr": "French", "hi": "Hindi"}
	codes["it"] = "Italian"

	fmt.Println(codes)
	delete(codes, "en")
	fmt.Println(codes)

}
