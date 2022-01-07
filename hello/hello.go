package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Declarando variáveis - forma 1
	// var nome string = "Romário"
	// var idade int = 26
	// var versao float32 = 1.1

	// Declarando variáveis - forma 2
	nome := "Romário"
	idade := 26
	versao := 1.1

	fmt.Println("Hello Mr.", nome,"! Sua idade é:", idade)
	fmt.Println("Nome é do tipo: ", reflect.TypeOf(nome))
	fmt.Println("Idade é do tipo: ", reflect.TypeOf(idade))
	fmt.Println("Versao é do tipo: ", reflect.TypeOf(versao))

	fmt.Println("\n\nVersão do software:", versao)
}