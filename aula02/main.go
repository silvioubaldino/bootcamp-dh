package main

import (
	".bootcamp-dh/aula02/bases"
	"fmt"
)

func main() {
	fmt.Println("-------- Aula 02 --------\n")
	fmt.Println("-------- Manhã --------\n")
	bases.CalculaImposto(50000)
	bases.CalculaImposto(150000)

	bases.Media(6, 5, 8, 7, 4)

	bases.CalculaSalario("C", 162)
	bases.CalculaSalario("B", 176)
	bases.CalculaSalario("A", 172)

	bases.ExecutaOperacao("maximo", 4, 5, 6, 7, 8)
	bases.ExecutaOperacao("media", 4, 5, 6, 7, 8)
	bases.ExecutaOperacao("minimo", 4, 5, 6, 7, 8)

	fmt.Println("-------- Tarde --------\n")
	estudante := bases.NewEstudante("Silvio", "Neto", "1212121", "12/04/2015")
	bases.ImprimeEstudante(estudante)

	bases.ImprimeEcommerce()
}
