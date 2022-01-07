package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	for {
		exibeMenu()

		comando := lerComando()

		switch comando {
			case 1:
				monitorar()
			case 2:
				exibeLogs()
			case 0:
				fmt.Println("Saindo do programa.")
				os.Exit(0)
			default:
				fmt.Println("Comando não reconhecido.")
				os.Exit(-1)
		}
	}
}

func exibeMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")
}

func lerComando() int {
	var comando int
	_, err := fmt.Scan(&comando)

	if err != nil {
		return -1
	}

	return comando
}

func monitorar() {
	fmt.Println("Informe uma url:")

	var url string
	fmt.Scan(&url)

	fmt.Println("Monitorando...")

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("===========")
		fmt.Println("Erro:")
		fmt.Println(err)
		fmt.Println("===========")
		return
	}

	fmt.Println("===========")
	fmt.Println(resp.Status)
	fmt.Println("===========")
}

func exibeLogs() {
	fmt.Println("Exibindo logs...")
}