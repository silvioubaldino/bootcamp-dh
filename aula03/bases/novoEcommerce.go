package bases

import "fmt"

func NovoEcommerce() {
	fmt.Println("-------- Exercício 02 --------")

	usuario1 := usuario{
		nome:      "Nome",
		sobrenome: "Sobrenome",
		email:     "nome@sobrenome.com",
	}

	p1 := novoProduto("Bicicleta", 2000)
	p2 := novoProduto("Patinete", 500)
	p3 := novoProduto("Patins", 2000)

	adicionarProduto(&usuario1, p1, 3)
	adicionarProduto(&usuario1, p2, 1)
	adicionarProduto(&usuario1, p3, 2)
	fmt.Println(usuario1)

	deletaProduto(&usuario1)

	fmt.Println(usuario1)

	fmt.Println()
}

type usuario struct {
	nome      string
	sobrenome string
	email     string
	produtos  []produto
}

type produto struct {
	nome  string
	preco float64
	qtd   int
}

func novoProduto(nome string, preco float64) produto {
	return produto{
		nome:  nome,
		preco: preco,
	}
}

func adicionarProduto(u *usuario, produto produto, qtd int) {
	produto.qtd = qtd
	u.produtos = append(u.produtos, produto)
}

func deletaProduto(u *usuario) {
	u.produtos = nil
}
