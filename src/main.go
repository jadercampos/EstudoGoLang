package main

import (
	"fmt"

	"estudogolang.com/src/clientes"
	"estudogolang.com/src/contas"
)

func main() {
	contaDoBruno := contas.ContaCorrente{Titular: clientes.Titular{
		Nome:      "Bruno",
		CPF:       "123.111.123.12",
		Profissao: "Desenvolvedor"},
		NumeroAgencia: 123, NumeroConta: 123456}
	fmt.Println(contaDoBruno)

	clienteSilvia := clientes.Titular{"Silvia", "321.546.789.44", "Puta"}
	contaDaSilvia := contas.ContaCorrente{Titular: clienteSilvia, NumeroAgencia: 456, NumeroConta: 789789}
	fmt.Println(contaDaSilvia)

}
