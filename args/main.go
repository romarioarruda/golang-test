package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	for {
		exibeMenu()

		comando := lerComando()

		switch comando {
		case 1:
			fmt.Println("Monitorando...")
			lerArquivosETestaSite()
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

func lerArquivosETestaSite() {
	arquivo, err := os.Open("sites.txt")

	if err != nil {
		log.Fatalln(err)
	}

	ponteiro := bufio.NewReader(arquivo)

	for {
		linha, err := ponteiro.ReadString('\n')
		linha = strings.TrimSpace(linha)

		fmt.Println("Testando site:", linha)
		testaSite(linha)
		fmt.Println("")

		if err == io.EOF {
			break
		}
	}

	arquivo.Close()
}

func testaSite(url string) {
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
