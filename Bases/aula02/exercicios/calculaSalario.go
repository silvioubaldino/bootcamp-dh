package exercicios

import "fmt"

func CalculaSalario(categoria string, horasMensais float32) {
	fmt.Println("-------- Exercício 03 --------")
	switch categoria {
	case "A":
		salario := 3000 * horasMensais
		if horasMensais > 160 {
			salario *= 1.5
		}
		fmt.Println(salario)
	case "B":
		salario := 1500 * horasMensais
		if horasMensais > 160 {
			salario *= 1.5
		}
		fmt.Println(salario)
	case "C":
		fmt.Println(1000 * horasMensais)
	default:
		fmt.Println("Categoria não encontrada")
	}
	fmt.Println()
}
