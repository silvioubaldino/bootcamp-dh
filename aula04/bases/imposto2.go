package bases

import "errors"

func VerificaSalario2(salario int) (string, error) {
	if salario < 15000 {
		return "", errors.New("erro: O salário digitado não está dentro do valor mínimo")
	}
	return "Necessário pagamento de imposto", nil
}
