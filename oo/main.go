package main

import (
	"fmt"
	"go-alura/oo/contas"
)

func main() {
	contaTati := contas.ContaCorrente{Titular: "Tati", Agencia: 2040, Conta: 999666, Saldo: 100}
	contaRomario := contas.ContaCorrente{Titular: "Romário", Agencia: 123, Conta: 45678, Saldo: 50}

	status := contaTati.Transferencia(&contaRomario, 90)

	fmt.Println(contaTati)
	fmt.Println(status)
	fmt.Println(contaRomario)
}