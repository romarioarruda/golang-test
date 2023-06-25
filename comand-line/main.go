package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
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

		registraLogs(url, "500 " + err.Error())
		return
	}

	fmt.Println("===========")
	fmt.Println(resp.Status)
	fmt.Println("===========")
	registraLogs(url, resp.Status)
}

func registraLogs(site string, statusCode string) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("===========")
		fmt.Println("Falha:")
		fmt.Println(err)
		fmt.Println("===========")
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - status=" + statusCode + "\n")

	fmt.Println("Log registrado")
	fmt.Println("===========")

	arquivo.Close()
}

func exibeLogs() {
	//Essa função abre, ler e fecha o arquivo
	//Diferente da função os.Open e os.OpenFile que precisa ser fechada manualmente.
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("======LOGS======")
	fmt.Println(string(arquivo))
	fmt.Println("======FIM LOGS======")
	fmt.Println("")
}
