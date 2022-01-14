package main

import (
	"fmt"
	"go-alura/oo/contas"
	// "go-alura/oo/clientes"
)

func main() {
	contaCliente := contas.ContaCorrente{}

	contaCliente.Deposito(100)

	fmt.Println(contaCliente.Saldo())
}