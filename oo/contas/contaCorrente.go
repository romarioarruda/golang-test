package contas

type ContaCorrente struct {
	Titular string
	Agencia int
	Conta int
	Saldo float64
}

func (c *ContaCorrente) Sacar(valor float64) string {
	podeSacar := valor > 0 && valor <= c.Saldo
	if podeSacar {
		c.Saldo -= valor

		return "Saque realizado com sucesso!"
	}

	return "Saldo insuficiente!"
}

func (c *ContaCorrente) Deposito(valor float64) string {
	podeDepositar := valor > 0
	if podeDepositar {
		c.Saldo += valor

		return "Depósito realizado com sucesso!"
	}

	return "Valor insuficiente para depósito!"
}

func (contaSaida *ContaCorrente) Transferencia(contaDestino *ContaCorrente, valor float64) string {
	if valor <= contaSaida.Saldo && valor > 0 {
		contaSaida.Sacar(valor)
		contaDestino.Deposito(valor)

		return "Transferência realizada com sucesso!"
	}

	return "Saldo insuficiente ou o valor é 0!"
}