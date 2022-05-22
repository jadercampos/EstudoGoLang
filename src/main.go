package main

import (
	"fmt"

	"estudogolang.com/src/pessoa"
)

func main() {
	p1 := pessoa.Pessoa{Nome: "João"}
	fmt.Println("Olá", p1.Nome)
}
