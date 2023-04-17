package main

func main() {

	var grades [5]int = [5]int{90, 80, 70, 80, 97}

	for index, element := range grades {
		println(index, "=>", element)
	}

}
