package main

import (
	bases2 ".bootcamp-dh/Bases/aula01/exercicios"
	"fmt"
)

func main() {
	fmt.Println("-------- Aula 01 --------\n\n")
	fmt.Println("-------- Manhã --------\n")
	bases2.PrintNomeIdade()
	bases2.PrintMeteorologia()
	bases2.CorrigirVariaveis()
	bases2.CorrigirTiposDeDados()
	fmt.Println("\n\n-------- Tarde --------\n")
	bases2.PrintLetras("Palavra")
	bases2.GetEmprestimo(21, true, 25, 3000)
	bases2.GetEmprestimo(23, false, 25, 3000)
	bases2.GetEmprestimo(23, true, 8, 3000)
	bases2.GetEmprestimo(23, true, 25, 3000)
	bases2.Mes(7)
	bases2.IdadePorNome("Benjamin")
}
