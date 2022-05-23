package bases

import "fmt"

func Media(notas ...float32) {
	fmt.Println("-------- Exercício 02 --------")
	var sum float32
	for _, nota := range notas {
		sum += nota
	}
	fmt.Println(sum / float32(len(notas)))
	fmt.Println()
}
