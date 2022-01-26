package main

import (
	"net/http"
	"go-alura/web-app/routes"
)


func main() {
	routes.CarregaRotasView()
	routes.CarregaRotasController()

	http.ListenAndServe(":8000", nil)
}
