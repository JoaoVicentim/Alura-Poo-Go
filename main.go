package main

import (
	"fmt"

	"github.com/JoaoVicentim/Alura-Poo-Go.git/contas"
)

func main() {
	// contaDoJoao := ContaCorrente{titular: "Joao",
	// 	numeroAgencia: 259, numeroConta: 123456, saldo: 125.6}
	// fmt.Println(contaDoJoao)

	// contaDaBruna := ContaCorrente{"Bruna", 222, 111222, 3029.4}
	// fmt.Println(contaDaBruna)

	// //*(ponteiro) ContaCorrente serve para indicar que estamos utilizando uma determinada ContaCorrente
	// //ponteiros "agem de baixo dos panos"
	// var contaDaGigi *ContaCorrente
	// contaDaGigi = new(ContaCorrente)
	// contaDaGigi.titular = "Gigi"
	// contaDaGigi.saldo = 328367

	// fmt.Println(*contaDaGigi)

	contaMarcela := contas.ContaCorrente{}
	contaMarcela.titular = "Marcela"
	contaMarcela.saldo = 1000

	contadaSilvia := contas.ContaCorrente{}
	contadaSilvia.titular = "Silvia"
	contadaSilvia.saldo = 500

	status := contaMarcela.Transferir(100, &contadaSilvia)
	fmt.Println(status)

	fmt.Println(contadaSilvia)
	fmt.Println(contaMarcela)

	//fmt.Println(contadaSilvia.Sacar(1000))
}
