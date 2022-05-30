package exercicios

import "fmt"

func ExecutaOperacao(operacao string, notas ...float32) {
	fmt.Println("-------- Exercício 04 --------")
	funcOperacao := getOperacao(operacao)
	fmt.Println(funcOperacao(notas))
	fmt.Println()
}

func getOperacao(operacao string) func(notas []float32) float32 {
	switch operacao {
	case "minimo":
		return minimo
	case "media":
		return media
	case "maximo":
		return maximo
	default:
		fmt.Println("Operação não encontrada")
		return nil
	}
}

func minimo(notas []float32) float32 {
	var min float32 = notas[0]
	for _, n := range notas {
		if n < min {
			min = n
		}
	}
	return min
}

func media(notas []float32) float32 {
	var sum float32
	for _, nota := range notas {
		sum += nota
	}
	return sum / float32(len(notas))
}

func maximo(notas []float32) float32 {
	var max float32 = notas[0]
	for _, n := range notas {
		if n > max {
			max = n
		}
	}
	return max
}
