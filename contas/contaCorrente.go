package contas

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

// Recebe como parâmetro o valorDoSaque
// Deve retornar uma string
// c * ContaCorrente -> Isso faz com que a função aponte para a conta que esta querendo realizar o saque
func (c *ContaCorrente) Sacar(valorDoSaque float64) (string, float64) {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.Saldo
	if podeSacar {
		c.Saldo -= valorDoSaque
		return "Saque realizado com sucesso!\nNovo Saldo disponível:", c.Saldo
	} else {
		return "Sem Saldo disponível!\nSaldo continua com:", c.Saldo
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	podeDepositar := valorDoDeposito > 0
	if podeDepositar {
		c.Saldo += valorDoDeposito
		return "Depósito realizado. Valor depositado:", valorDoDeposito
	} else {
		return "Depósito não realizado", 0
	}

}

func (c *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) bool {

	if valorTransferencia < c.Saldo && valorTransferencia > 0 {
		c.Saldo -= valorTransferencia

		contaDestino.Depositar(valorTransferencia)
		return true
	} else {
		return false
	}
}
