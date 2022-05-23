package bases

import "fmt"

func PrintLetras(palavra string) {
	fmt.Println("-------- Exercício 01 --------")

	fmt.Println(len(palavra))

	for _, letra := range palavra {
		fmt.Printf("%c, ", letra)
	}

	fmt.Print("\n\n")
}
