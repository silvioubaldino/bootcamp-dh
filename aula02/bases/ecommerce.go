package bases

type loja struct {
	produtos []produto
}

func (l loja) Total() {
	//preco do produto + custos
}

func (l *loja) Adicionar(p produto) {
	l.produtos = append(l.produtos, p)
}

func NovoEcommerce() Ecommerce {
	//case "TipoEcommerce" == loja
	return &loja{}
}

type Ecommerce interface {
	Total()
	Adicionar(p produto)
}

type produto struct {
	tipoDeProduto string
	nome          string
	preco         float64
}

func NovoProduto(tipoDeProduto string, nome string, preco float64) produto {
	return produto{
		tipoDeProduto: tipoDeProduto,
		nome:          nome,
		preco:         preco,
	}
}

type Produto interface {
	CalculaCusto()
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
