package main

import (
	"fmt"
	"objetos/contas"
)

func main() {
	contaDoPedro := contas.ContaCorrente{}
	contaDoPedro.titular = "Pedro"
	contaDoPedro.numeroAgencia = 1234
	contaDoPedro.numeroConta = 654321
	contaDoPedro.saldo = 1587.45
	fmt.Println(contaDoPedro)

	contaDoNicolas := contas.ContaCorrente{titular: "Nicolas", saldo: 500}

	valorDoSaque := 200
	contaDoPedro.saldo = contaDoPedro.saldo - float64(valorDoSaque)
	fmt.Println(contaDoPedro.saldo)

	fmt.Println(contaDoPedro.Sacar(300))
	fmt.Println(contaDoPedro.saldo)
	fmt.Println(contaDoPedro.Sacar(3000))
	fmt.Println(contaDoPedro.saldo)

	fmt.Println(contaDoPedro.Depositar(5000))
	status, valor := contaDoPedro.Depositar(-5000)
	fmt.Println(status, " - Saldo: ", valor)

	retorno := contaDoPedro.Transferir(100, &contaDoNicolas)
	fmt.Println(retorno)
}
