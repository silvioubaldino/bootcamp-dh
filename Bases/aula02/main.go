package main

import (
	bases2 ".bootcamp-dh/Bases/aula02/exercicios"
	"fmt"
)

func main() {
	fmt.Println("-------- Aula 02 --------\n")
	fmt.Println("-------- Manhã --------\n")
	bases2.CalculaImposto(50000)
	bases2.CalculaImposto(150000)

	bases2.Media(6, 5, 8, 7, 4)

	bases2.CalculaSalario("C", 162)
	bases2.CalculaSalario("B", 176)
	bases2.CalculaSalario("A", 172)

	bases2.ExecutaOperacao("maximo", 4, 5, 6, 7, 8)
	bases2.ExecutaOperacao("media", 4, 5, 6, 7, 8)
	bases2.ExecutaOperacao("minimo", 4, 5, 6, 7, 8)

	fmt.Println("-------- Tarde --------\n")
	estudante := bases2.NewEstudante("Silvio", "Neto", "1212121", "12/04/2015")
	bases2.ImprimeEstudante(estudante)

	bases2.ImprimeEcommerce()
}
