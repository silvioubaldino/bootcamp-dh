package bases

type error interface {
	Error() string
}

type salarioError struct {
	msg string
}

func (e *salarioError) Error() string {
	return e.msg
}

func VerificaSalario(salario int) (string, error) {
	if salario < 15000 {
		return "", &salarioError{msg: "erro: O salário digitado não está dentro do valor mínimo"}
	}
	return "Necessário pagamento de imposto", nil
}
