package models

import (
	"go-alura/web-app/db"
)

type Produto struct {
	Id, Quantidade int
	Nome, Descricao, Preco string
}

func GetProdutos() []Produto {
	db := db.ConectaBancoDeDados()

	selectProdutos, err := db.Query("SELECT * FROM produtos")

	if err != nil {
		panic(err.Error())
	}

	produto := Produto{}
	produtos := []Produto{}

	var id, quantidade int
	var nome, descricao, preco string

	for selectProdutos.Next() {
		err := selectProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		produto.Id = id
		produto.Nome = nome
		produto.Descricao = descricao
		produto.Preco = preco
		produto.Quantidade = quantidade

		produtos = append(produtos, produto)
	}

	defer db.Close()

	return produtos
}
