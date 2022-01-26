package routes

import (
	"net/http"
	"go-alura/web-app/controllers"
)


func CarregaRotasView() {
	http.HandleFunc("/", controllers.ListaTodosOsProdutos)
}

func CarregaRotasController() {
	http.HandleFunc("/insert", controllers.CriaNovoProduto)
}
