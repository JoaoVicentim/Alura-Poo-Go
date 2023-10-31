package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

//Recebe como parâmetro o valorDoSaque
//Deve retornar uma string
//c * ContaCorrente -> Isso faz com que a função aponte para a conta que esta querendo realizar o saque
func (c *ContaCorrente) Sacar(valorDoSaque float64) (string, float64) {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso!\nNovo saldo disponível:", c.saldo
	} else {
		return "Sem saldo disponível!\nSaldo continua com:", c.saldo
	}
}

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

	contadaSilvia := ContaCorrente{}
	contadaSilvia.titular = "Silvia"
	contadaSilvia.saldo = 500

	fmt.Println(contadaSilvia)

	fmt.Println(contadaSilvia.Sacar(1000))
	fmt.Println("spinola")
}
