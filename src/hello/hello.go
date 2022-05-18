package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const monitoramentos, delay = 3, 2

func main() {
	exibeSaudacao()
	for {
		exibeMenu()
		executaComando()
	}

}
func devolveNomeEIdade() (string, int) {
	var nome string = "Jader"
	var idade = 39
	return nome, idade
}
func exibeSaudacao() {
	nome, idade := devolveNomeEIdade()
	versao := 1.2
	fmt.Println("Olá, sr.", nome, "sua idade é", idade)
	fmt.Println("Meu primeiro programa em go, versão", versao)
	fmt.Println("O tipo da variável nome é ", reflect.TypeOf(nome))
}
func exibeMenu() {
	fmt.Println("\n---// OPÇÕES //---:")
	fmt.Println("\n1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("3- Sair do Programa")
}
func leComando() int {
	var comandoLido int
	fmt.Print("\nDigite a sua opção: ")
	fmt.Scan(&comandoLido)
	fmt.Printf("\nO comando escolhido foi %d (endereço: %d)\n", comandoLido, &comandoLido)
	return comandoLido
}
func executaComando() {
	comando := leComando()
	switch comando {
	case 1:
		iniciarMonitoramento()
	case 2:
		imprimeLogs()
	case 3:
		fmt.Println("\nSaindo ...")
		os.Exit(0)
	default:
		fmt.Println("\nNão conheço este comando")
		os.Exit(-1)
	}
}
func iniciarMonitoramento() {
	fmt.Println("\nMonitorando ...")
	sites := leSitesDoArquivo()
	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site ", i+1, " : ", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println(" ** // ... // ** ")
	}
}
func testaSite(site string) {
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
		return
	}
	if resp.StatusCode == 200 {
		fmt.Println("O site", site, " foi carregaado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("O site", site, " está com problemas. Status Code: ", resp.StatusCode)
		registraLog(site, false)
	}

}
func leSitesDoArquivo() []string {
	var sites []string
	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
		return sites
	}
	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
		if err == io.EOF {
			break
		}
	}
	arquivo.Close()
	return sites
}
func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
		return
	}
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")
	arquivo.Close()
}
func imprimeLogs() {
	fmt.Println("\nExibindo Logs  ...")
	arquivo, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
		return
	}
	fmt.Println(string(arquivo))
}
