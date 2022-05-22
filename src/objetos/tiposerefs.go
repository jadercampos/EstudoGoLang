package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {

	contaDoJader := ContaCorrente{
		titular:       "Jader",
		numeroAgencia: 589,
		numeroConta:   123456,
		saldo:         125.50}
	fmt.Println(contaDoJader)

	contaDaDri := ContaCorrente{"Adriana", 123, 654, 45.45}
	fmt.Println(contaDaDri)

	fmt.Println(ContaCorrente{})

	var titular string = "Jader"
	var numeroAgencia int = 589
	var numeroConta int = 123456
	var saldo float64 = 125.50
	fmt.Println(titular, numeroAgencia, numeroConta, saldo)

	var titular2 string = "Adriana"
	var numeroAgencia2 int = 487
	var numeroConta2 int = 654321
	var saldo2 float64 = 100.58
	fmt.Println(titular2, numeroAgencia2, numeroConta2, saldo2)

	var contaDoJoao *ContaCorrente
	contaDoJoao = new(ContaCorrente)
	contaDoJoao.titular = "João"
	contaDoJoao.numeroAgencia = 528
	contaDoJoao.numeroConta = 785421
	contaDoJoao.saldo = 500

	fmt.Println(contaDoJoao)

	contaDaDri := ContaCorrente{"Adriana", 123, 654, 45.45}
	contaDaDri2 := ContaCorrente{"Adriana", 123, 654, 45.45}
	fmt.Println(contaDaDri == contaDaDri2)

	var contaDoJoao *ContaCorrente
	contaDoJoao = new(ContaCorrente)
	contaDoJoao.titular = "João"
	contaDoJoao.numeroAgencia = 528
	contaDoJoao.numeroConta = 785421
	contaDoJoao.saldo = 500

	var contaDoJoao2 *ContaCorrente
	contaDoJoao2 = new(ContaCorrente)
	contaDoJoao2.titular = "João"
	contaDoJoao2.numeroAgencia = 528
	contaDoJoao2.numeroConta = 785421
	contaDoJoao2.saldo = 500

	fmt.Println(contaDoJoao == contaDoJoao2)
	fmt.Println(contaDoJoao)
	fmt.Println(contaDoJoao2)
	fmt.Println(&contaDoJoao)
	fmt.Println(&contaDoJoao2)
}
