package controllers

import (
	"go-alura/web-app/models"
	"html/template"
	"net/http"
)


var temp = template.Must(template.ParseGlob("templates/*.html"))

func ListaTodosOsProdutos(resp http.ResponseWriter, req *http.Request) {
	produtos := models.GetProdutos()

	temp.ExecuteTemplate(resp, "Index", produtos)
}
