package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)


func conectaBancoDeDados()  *sql.DB {
	conexao := "user=postgres dbname=postgres password=root sslmode=disable"

	db, err := sql.Open("postgres", conexao)

	if err != nil {
		panic(err.Error())
	}

	return db
}

type Produto struct {
	Id, Quantidade int
	Nome, Descricao, Preco string
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(resp http.ResponseWriter, req *http.Request) {
	db := conectaBancoDeDados()

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

	temp.ExecuteTemplate(resp, "Index", produtos)
	defer db.Close()
}
