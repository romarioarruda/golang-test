package contas

import (
	"go-alura/oo/clientes"
)

type ContaCorrente struct {
	Titular clientes.Titular
	Agencia int
	Conta int
	saldo float64
}

func (c *ContaCorrente) Saldo() (string, float64) {
	return "Seu saldo é de: R$", c.saldo
}

func (c *ContaCorrente) Sacar(valor float64) string {
	podeSacar := valor > 0 && valor <= c.saldo
	if podeSacar {
		c.saldo -= valor

		return "Saque realizado com sucesso!"
	}

	return "Saldo insuficiente!"
}

func (c *ContaCorrente) Deposito(valor float64) string {
	podeDepositar := valor > 0
	if podeDepositar {
		c.saldo += valor

		return "Depósito realizado com sucesso!"
	}

	return "Valor insuficiente para depósito!"
}

func (contaSaida *ContaCorrente) Transferencia(contaDestino *ContaCorrente, valor float64) string {
	if valor <= contaSaida.saldo && valor > 0 {
		contaSaida.Sacar(valor)
		contaDestino.Deposito(valor)

		return "Transferência realizada com sucesso!"
	}

	return "Saldo insuficiente ou o valor é 0!"
}