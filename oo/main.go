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

func (contaSaida *ContaCorrente) transferencia(contaDestino *ContaCorrente, valor float64) string {
	if valor <= contaSaida.saldo && valor > 0 {
		contaSaida.sacar(valor)
		contaDestino.deposito(valor)

		return "Transferência realizada com sucesso!"
	}

	return "Saldo insuficiente ou o valor é 0!"
}


func main() {
	contaTati := ContaCorrente{"Tati", 2040, 999666, 100}
	contaRomario := ContaCorrente{"Romario", 123, 45678, 100}

	status := contaTati.transferencia(&contaRomario, 90)

	fmt.Println(contaTati)
	fmt.Println(status)
	fmt.Println(contaRomario)
	
}