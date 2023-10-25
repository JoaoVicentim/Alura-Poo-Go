package main

import (
	"fmt"
)

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaDoJoao := ContaCorrente{titular: "Joao",
		numeroAgencia: 259, numeroConta: 123456, saldo: 125.6}
	fmt.Println(contaDoJoao)

	contaDaBruna := ContaCorrente{"Bruna", 222, 111222, 3029.4}
	fmt.Println(contaDaBruna)

	//*(ponteiro) ContaCorrente serve para indicar que estamos utilizando uma determinada ContaCorrente
	var contaDaGigi *ContaCorrente
	contaDaGigi = new(ContaCorrente)
	contaDaGigi.titular = "Gigi"
	contaDaGigi.saldo = 328367

	fmt.Println(contaDaGigi)

}