package routes

import (
	"net/http"
	"go-alura/web-app/controllers"
)


func CarregaRotas() {
	http.HandleFunc("/", controllers.ListaTodosOsProdutos)
}
