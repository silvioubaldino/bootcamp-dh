package bases

import (
	"fmt"
)

func VerificaSalario3(salario int) (string, error) {
	if salario < 15000 {
		return "", fmt.Errorf("erro: o mínimo tributável é 15.000 e o salário informado é: %d", salario)
	}
	return "Necessário pagamento de imposto", nil
}
