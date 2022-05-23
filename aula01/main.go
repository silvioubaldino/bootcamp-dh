package main

import (
	"fmt"

	".bootcamp-dh/aula01/bases"
)

func main() {
	fmt.Println("-------- Aula 01 --------\n\n")
	fmt.Println("-------- Manhã --------\n")
	bases.PrintNomeIdade()
	bases.PrintMeteorologia()
	bases.CorrigirVariaveis()
	bases.CorrigirTiposDeDados()
	fmt.Println("\n\n-------- Tarde --------\n")
	bases.PrintLetras("Palavra")
	bases.GetEmprestimo(21, true, 25, 3000)
	bases.GetEmprestimo(23, false, 25, 3000)
	bases.GetEmprestimo(23, true, 8, 3000)
	bases.GetEmprestimo(23, true, 25, 3000)
	bases.Mes(7)
	bases.IdadePorNome("Benjamin")
}
