package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque sucesso"
	} else {
		return "Saldo Insuficiente!"
	}
}

func main() {
	contaDoPedro := ContaCorrente{}
	contaDoPedro.titular = "Pedro"
	contaDoPedro.numeroAgencia = 1234
	contaDoPedro.numeroConta = 654321
	contaDoPedro.saldo = 1587.45
	fmt.Println(contaDoPedro)

	valorDoSaque := 200
	contaDoPedro.saldo = contaDoPedro.saldo - float64(valorDoSaque)
	fmt.Println(contaDoPedro.saldo)

	fmt.Println(contaDoPedro.Sacar(300))
	fmt.Println(contaDoPedro.saldo)
	fmt.Println(contaDoPedro.Sacar(3000))
	fmt.Println(contaDoPedro.saldo)
}
