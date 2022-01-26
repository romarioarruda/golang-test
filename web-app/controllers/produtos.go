package controllers

import (
	"go-alura/web-app/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)


var temp = template.Must(template.ParseGlob("views/*.html"))

func ListaTodosOsProdutos(resp http.ResponseWriter, req *http.Request) {
	produtos := models.GetProdutos()

	temp.ExecuteTemplate(resp, "Index", produtos)
}

func CriaNovoProduto(resp http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		nome := req.FormValue("nome")
		descricao := req.FormValue("descricao")
		preco := req.FormValue("preco")
		quantidade := req.FormValue("quantidade")

		precoConvertidoParaFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}

		quantidadeConvertidaParaInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão da quantidade:", err)
		}

		models.CriaNovoProduto(nome, descricao, precoConvertidoParaFloat, quantidadeConvertidaParaInt)
	}

	http.Redirect(resp, req, "/", http.StatusMovedPermanently)
}
