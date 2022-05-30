package exercicios

import "fmt"

type loja struct {
	produtos []Produto
}

func (l loja) Total() float64 {
	var total float64
	for _, p := range l.produtos {
		total += p.GetPreco() + p.CalcularCusto()
	}
	return total
}

func (l *loja) Adicionar(produtos ...Produto) {
	for _, p := range produtos {
		l.produtos = append(l.produtos, p)
	}
}

func NovoEcommerce() Ecommerce {
	//case "TipoEcommerce" == loja
	return &loja{}
}

type Ecommerce interface {
	Total() float64
	Adicionar(produtos ...Produto)
}

type produto struct {
	tipoDeProduto string
	nome          string
	preco         float64
}

type Produto interface {
	CalcularCusto() float64
	GetPreco() float64
}

func NovoProduto(tipoDeProduto string, nome string, preco float64) *produto {
	return &produto{
		tipoDeProduto: tipoDeProduto,
		nome:          nome,
		preco:         preco,
	}
}

func (p produto) CalcularCusto() float64 {
	switch p.tipoDeProduto {
	case "pequeno":
		return 0
	case "medio":
		return p.preco * 0.03
	case "grande":
		return p.preco*0.06 + 2500
	default:
		return 0
	}
}

func (p produto) GetPreco() float64 {
	return p.preco
}

func ImprimeEcommerce() {
	fmt.Println("-------- Exercício 02 --------")
	loja01 := NovoEcommerce()
	produto01 := NovoProduto("pequeno", "Celular", 5000)
	produto02 := NovoProduto("medio", "TV", 5000)
	produto03 := NovoProduto("grande", "Geladeira", 5000)
	loja01.Adicionar(produto01, produto02, produto03)
	fmt.Println("Loja 01")
	fmt.Println(produto01.CalcularCusto())
	fmt.Println(produto02.CalcularCusto())
	fmt.Println(produto03.CalcularCusto())
	fmt.Println(loja01.Total())
	fmt.Println()

	lojaA := NovoEcommerce()
	produtoA := NovoProduto("pequeno", "Chocolate", 10)
	produtoB := NovoProduto("medio", "Leite", 10)
	produtoC := NovoProduto("grande", "Vassoura", 10)
	lojaA.Adicionar(produtoA, produtoB, produtoC)
	fmt.Println("Loja A")
	fmt.Println(produtoA.CalcularCusto())
	fmt.Println(produtoB.CalcularCusto())
	fmt.Println(produtoC.CalcularCusto())
	fmt.Println(lojaA.Total())
	fmt.Println()
}
