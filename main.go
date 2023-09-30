package main

import (
	"fmt"

	cl "github.com/adrielldev/bank-go/clientes"
	c "github.com/adrielldev/bank-go/contas"
)

type verificarConta interface {
	Sacar(valor float64) string
}

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

func main() {
	clienteAdriel := cl.Titular{
		Nome:      "Adriel",
		CPF:       "123456789",
		Profissao: "Desenvolvedor",
	}

	contaDoAdriel := c.ContaCorrente{Titular: clienteAdriel,
		NumeroAgencia: 123,
		NumeroConta:   1,
	}

	contaDoAdriel.Depositar(500)
	contaPoupanca := c.ContaPoupanca{Titular: clienteAdriel, NumeroAgencia: 123, NumeroConta: 1, Operacao: 2}
	PagarBoleto(&contaPoupanca, 300)
	fmt.Println(contaPoupanca.ObterSaldo())
}
