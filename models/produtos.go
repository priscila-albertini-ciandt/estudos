package models

import (
	"alura/db"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosOsProdutos() []Produto {
	db := db.ConectaComBancoDeDados()

	selectProdutos, erro := db.Query("SELECT * FROM produtos ORDER BY id")

	if erro != nil {
		panic(erro.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		erro = selectProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if erro != nil {
			panic(erro.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	defer db.Close()

	return produtos
}

func CriarNovoProduto(nome, descricao string, preco float64, quantidade int) {

	db := db.ConectaComBancoDeDados()

	insereDadosNoBanco, erro := db.Prepare("insert into produtos (nome, descricao, preco, quantidade) values ($1, $2, $3, $4)")

	if erro != nil {
		panic(erro.Error())
	}

	insereDadosNoBanco.Exec(nome, descricao, preco, quantidade)

	defer db.Close()
}

func DeletarProduto(id string) {

	db := db.ConectaComBancoDeDados()

	deletaDoBanco, erro := db.Prepare("delete from produtos where id = $1")

	if erro != nil {
		panic(erro.Error())
	}

	deletaDoBanco.Exec(id)

	defer db.Close()
}

func BuscaProdutoPorId(id string) Produto {

	db := db.ConectaComBancoDeDados()

	produtoDoBanco, erro := db.Query("select * from produtos where id = $1", id)

	if erro != nil {
		panic(erro.Error())
	}

	produto := Produto{}

	for produtoDoBanco.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		erro = produtoDoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if erro != nil {
			panic(erro.Error())
		}

		produto.Id = id
		produto.Nome = nome
		produto.Descricao = descricao
		produto.Preco = preco
		produto.Quantidade = quantidade
	}

	defer db.Close()

	return produto
}

func AtualizaProduto(id, quantidade int, nome, descricao string, preco float64) {

	db := db.ConectaComBancoDeDados()

	atualizaProduto, erro := db.Prepare("update produtos set nome = $1, descricao = $2, preco = $3, quantidade = $4 where id = $5")

	if erro != nil {
		panic(erro.Error())
	}

	atualizaProduto.Exec(nome, descricao, preco, quantidade, id)

	defer db.Close()

}
