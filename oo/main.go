package main

import "fmt"

type ContaCorrente struct {
	titular string
	agencia int
	conta int
	saldo float64
}


func main() {
	// forma 1
	conta1 := ContaCorrente{
		titular: "Romário",
		agencia: 4733,
		conta: 212407,
		saldo: 100,
	}

	// forma 2
	conta2 := ContaCorrente{"Tati", 2040, 999666,100}

	// forma 3
	conta3 := new(ContaCorrente)
	conta3.titular = "Maria"
	conta3.agencia = 123
	conta3.conta = 567
	conta3.saldo = 10.99

	fmt.Println(conta1)
	fmt.Println(conta2)
	fmt.Println(conta3)
}