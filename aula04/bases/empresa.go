package bases

import "errors"

func CalculaSalario(horasTrabalhadas float64, valorHora float64) (float64, error) {
	if horasTrabalhadas < 80 {
		return 0, errors.New("erro: o trabalhador não pode ter trabalhado menos de 80 horas por mês")
	}
	salarioMensal := horasTrabalhadas * valorHora
	if salarioMensal >= 20000 {
		salarioMensal -= salarioMensal * 0.1
	}
	return salarioMensal, nil
}
