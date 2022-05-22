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

	contaDaDri2 := ContaCorrente{"Adriana", 123, 654, 45.45}
	contaDaDri3 := ContaCorrente{"Adriana", 123, 654, 45.45}
	fmt.Println(contaDaDri2 == contaDaDri3)

	var contaDoJoao2 *ContaCorrente
	contaDoJoao2 = new(ContaCorrente)
	contaDoJoao2.titular = "João"
	contaDoJoao2.numeroAgencia = 528
	contaDoJoao2.numeroConta = 785421
	contaDoJoao2.saldo = 500

	var contaDoJoao3 *ContaCorrente
	contaDoJoao3 = new(ContaCorrente)
	contaDoJoao3.titular = "João"
	contaDoJoao3.numeroAgencia = 528
	contaDoJoao3.numeroConta = 785421
	contaDoJoao3.saldo = 500

	fmt.Println(contaDoJoao2 == contaDoJoao3)
	fmt.Println(contaDoJoao2)
	fmt.Println(contaDoJoao3)
	fmt.Println(&contaDoJoao2)
	fmt.Println(&contaDoJoao3)
}
