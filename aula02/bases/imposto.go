package bases

import "fmt"

func CalculaImposto(salario float32) {
	fmt.Println("-------- Exercício 01 --------")

	if salario < 150000 {
		fmt.Println(salario * 0.17)
	} else {
		fmt.Println(salario * 0.27)
	}
	fmt.Println()
}
