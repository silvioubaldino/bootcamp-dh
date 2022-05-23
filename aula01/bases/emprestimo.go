package bases

import (
	"errors"
	"fmt"
)

func GetEmprestimo(idade uint8, empregado bool, mesesDeTrabalho uint8, salario float32) bool {
	//retorno nao especificado

	fmt.Println("-------- Exercício 02 --------")
	if err := IsAutorizadoEmprestimo(idade, empregado, mesesDeTrabalho); err != nil {
		fmt.Println(err)
		fmt.Println()
		return false
	}

	//CalculaJuros(salario)
	fmt.Println("Emprestimo aprovado")
	fmt.Println()
	return true
}

func CalculaJuros(salario float32) float32 {
	if salario > 100000 {
		return 0
	}
	return 0 //a ser implementados
}

func IsAutorizadoEmprestimo(idade uint8, empregado bool, mesesDeTrabalho uint8) error {
	if idade <= 22 {
		return errors.New("o cliente possui menos de 22 anos")
	} else if !empregado {
		return errors.New("o cliente não está empregado")
	} else if mesesDeTrabalho < 12 {
		return errors.New("o cliente tem menos de 12 meses de trabalho")
	}
	return nil
}
