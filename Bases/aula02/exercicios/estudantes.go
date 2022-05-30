package exercicios

import "fmt"

type estudante struct {
	nome      string
	sobrenome string
	rg        string
	dtEmissao string
}

func NewEstudante(nome string, sobrenome string, rg string, dtEmissao string) estudante {
	return estudante{
		nome:      nome,
		sobrenome: sobrenome,
		rg:        rg,
		dtEmissao: dtEmissao,
	}
}

func ImprimeEstudante(estudante estudante) {
	fmt.Println("-------- Exercício 01 --------")
	fmt.Println(estudante)
	fmt.Println()
}
