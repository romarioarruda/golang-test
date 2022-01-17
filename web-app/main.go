package main

import (
	"html/template"
	"net/http"
)

type Produto struct {
	Nome, Descricao string
	Preco float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(resp http.ResponseWriter, req *http.Request) {
	produtos := []Produto{
		{ "Caneca da Original.io", "Essa caneca é maneira", 22.99, 5 },
		{ "Cubo Mágico", "Um cubo muito massa", 50, 3 },
		{ "Celular", "Um celular bom zim", 900, 3 },
	}

	temp.ExecuteTemplate(resp, "Index", produtos)
}
