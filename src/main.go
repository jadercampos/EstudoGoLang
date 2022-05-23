package main

import (
	"fmt"

	"estudogolang.com/src/clientes"
	"estudogolang.com/src/contas"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaDoBruno := contas.ContaCorrente{Titular: clientes.Titular{
		Nome:      "Bruno",
		CPF:       "123.111.123.12",
		Profissao: "Desenvolvedor"},
		NumeroAgencia: 123, NumeroConta: 123456, Operacao: 1}
	contaDoBruno.Depositar(5000)
	PagarBoleto(&contaDoBruno, 500)
	fmt.Println(contaDoBruno)

	clienteSilvia := clientes.Titular{"Silvia", "321.546.789.44", "Puta"}
	contaDaSilvia := contas.ContaPoupanca{Titular: clienteSilvia, NumeroAgencia: 456, NumeroConta: 789789, Operacao: 13}
	contaDaSilvia.Depositar(2000)
	PagarBoleto(&contaDaSilvia, 500)
	fmt.Println(contaDaSilvia)

}
