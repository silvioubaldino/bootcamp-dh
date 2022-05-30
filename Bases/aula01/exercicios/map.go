package exercicios

import "fmt"

func IdadePorNome(nome string) {
	fmt.Println("-------- Exercício 04 --------")

	var employees = map[string]int{"Benjamin": 20, "Manuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}
	fmt.Printf("A idade de %s é: %d\n", nome, employees[nome])

	var count uint8
	for _, idade := range employees {
		if idade > 21 {
			count++
		}
	}
	fmt.Printf("%d pessoas tem mais de 21 anos\n", count)

	employees["Federico"] = 25

	fmt.Println(employees)

	delete(employees, "Pedro")

	fmt.Println(employees)

	fmt.Println()
}
