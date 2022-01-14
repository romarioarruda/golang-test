package main

import (
	"fmt"
	"go-alura/oo/contas"
	// "go-alura/oo/clientes"
)

type Conta interface {
	Sacar(valor float64) string
}

func PagarBoleto(conta Conta, valorDoBoleto float64) {
	status := conta.Sacar(valorDoBoleto)

	fmt.Println(status)
}

func main() {
	contaCliente := contas.ContaCorrente{}
	contaCliente.Deposito(100)

	PagarBoleto(&contaCliente, 60)

	fmt.Println(contaCliente.Saldo())
}