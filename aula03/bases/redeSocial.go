package bases

import (
	"fmt"
)

type cliente struct {
	nome      string
	sobrenome string
	idade     uint8
	email     string
	senha     string
}

func RedeSocial() {
	fmt.Println("-------- Exercício 01 --------")

	usuario1 := cliente{
		nome:      "Pedro",
		sobrenome: "Antonio",
		idade:     22,
		email:     "pedro@pedro.com",
		senha:     "pedropedro",
	}
	usuario2 := cliente{
		nome:      "Joao",
		sobrenome: "Antonio",
		idade:     26,
		email:     "joao@pedro.com",
		senha:     "joaojoao",
	}

	fmt.Printf("%s - %s - %d - %s - %s\n", usuario1.nome, usuario1.sobrenome, usuario1.idade, usuario1.email, usuario1.senha)
	fmt.Printf("%s - %s - %d - %s - %s\n\n", usuario2.nome, usuario2.sobrenome, usuario2.idade, usuario2.email, usuario2.senha)

	usuario1.SetNome("Carlos")
	usuario1.SetSobrenome("Alberto")
	usuario1.SetIdade(50)
	usuario1.SetEmail("carlos@carlos.br")
	usuario1.SetSenha("senhas132")

	usuario1.SetNome("Jorge")
	usuario1.SetSobrenome("Silva")
	usuario1.SetIdade(35)
	usuario1.SetEmail("jorge@manga.br")
	usuario2.SetSenha("senhas11111")

	fmt.Printf("%s - %s - %d - %s - %s\n", usuario1.nome, usuario1.sobrenome, usuario1.idade, usuario1.email, usuario1.senha)
	fmt.Printf("%s - %s - %d - %s - %s\n\n", usuario2.nome, usuario2.sobrenome, usuario2.idade, usuario2.email, usuario2.senha)
}

func (u *cliente) SetNome(novoNome string) {
	u.nome = novoNome
}

func (u *cliente) SetSobrenome(novoSobrenome string) {
	u.sobrenome = novoSobrenome
}
func (u *cliente) SetIdade(novaIdade uint8) {
	u.idade = novaIdade
}
func (u *cliente) SetEmail(novoEmail string) {
	u.email = novoEmail
}
func (u *cliente) SetSenha(novoSenha string) {
	u.senha = novoSenha
}
