package main

import (
	"net/http"
	"go-alura/web-app/routes"
)


func main() {
	routes.CarregaRotas()

	http.ListenAndServe(":8000", nil)
}
