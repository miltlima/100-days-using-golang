package main

import "fmt"

func main() {
	var materia1 float64
	var materia2 float64
	var materia3 float64
	var materia4 float64

	fmt.Println("Informe a nota em Materia 1: ")
	fmt.Scanf("%f", &materia1)
	fmt.Println("Informe a nota em Materia 2: ")
	fmt.Scanf("%f", &materia2)
	fmt.Println("Informe a nota em Materia 3: ")
	fmt.Scanf("%f", &materia3)
	fmt.Println("Informe a nota em Materia 4: ")
	fmt.Scanf("%f", &materia4)

	media := (materia1 + materia2 + materia3 + materia4) / 4

	if media >= 7 {
		fmt.Println("PArabéns, Sua média final foi:", media, "você foi aprovado")
	} else {
		fmt.Println("Você foi reprovado!")
	}

}
