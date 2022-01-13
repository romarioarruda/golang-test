package main

import "fmt"

type ContaCorrente struct {
	titular string
	agencia int
	conta int
	saldo float64
}

func (c *ContaCorrente) sacar(valor float64) string {
	podeSacar := valor > 0 && valor <= c.saldo
	if podeSacar {
		c.saldo -= valor

		return "Saque realizado com sucesso!"
	}

	return "Saldo insuficiente!"
}

func (c *ContaCorrente) deposito(valor float64) string {
	podeDepositar := valor > 0
	if podeDepositar {
		c.saldo += valor

		return "Depósito realizado com sucesso!"
	}

	return "Valor insuficiente para depósito!"
}


func main() {
	conta := ContaCorrente{"Tati", 2040, 999666, 100}
	conta.saldo = 500

	fmt.Println(conta)
	fmt.Println(conta.sacar((-400)))
	fmt.Println(conta)

	fmt.Println(conta.deposito((10)))
	fmt.Println(conta)
}