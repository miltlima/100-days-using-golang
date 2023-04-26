package main

import "fmt"

type Student  struct {
	name string
	grades []int

}

func (s *Student) DisplayName() {
	fmt.Println(s.name)
}

func (s *Student) calculatePercentage() float64 {
	sum := 0 
	for _, v := range s.grades {
		sum += v
	}
	return float64(sum*100) / float64(len(s.grades)*100)
}

func main() {
	s := Student{name: "Milton",grades: []int{90, 75, 80}}
	s.DisplayName()
	fmt.Printf("%.2f%%", s.calculatePercentage())
}