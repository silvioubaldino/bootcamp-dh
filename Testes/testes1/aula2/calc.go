package aula2

import "fmt"

func Sub(a, b int) int {
	return a - b
}

func Divide(num, den int) (int, error) {
	if den == 0 {
		return 0, fmt.Errorf("o denominador não pode ser 0")
	}
	return num / den, nil
}
