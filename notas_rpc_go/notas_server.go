package main

import (
	"fmt"
	"net"
	"net/rpc"
)

// Estrutura para representar um aluno
type Conta struct {
	Nome  string
	Saldo float64
}

// Estrutura para o servidor
type Servidor struct {
	contas []Conta
}

// Método para inicializar a lista de alunos no servidor
func (s *Servidor) inicializar() {
	s.contas = []Conta{
		{"Alexandre", 900.5},
		{"Barbara", 855.5},
		{"Joao", 656.5},
		{"Maria", 946.0},
		{"Paulo", 10546.0},
		{"Pedro", 7465.0},
	}
}

// Método remoto que retorna a nota de um aluno dado o seu nome
func (s *Servidor) ObtemSaldo(nome string, saldo *float64) error {
	for _, conta := range s.contas {
		if conta.Nome == nome {
			*saldo = conta.Saldo
			return nil
		}
	}
	return fmt.Errorf("Aluno %s não encontrado", nome)
}

func (s *Servidor) AbrirConta(conta Conta, resposta *string) error {

	if conta.Saldo < 0 {
		*resposta = "Erro saldo inicial invalido"

	} else {
		// Adiciona a nova conta à lista de contas
		s.contas = append(s.contas, conta)
		*resposta = "Conta criada com sucesso!"
	}
	return nil
}
func (s *Servidor) FecharConta(conta Conta, resposta *string) error {

	for i, a := range s.contas {
		if a.Nome == conta.Nome {
			// Remove a conta da lista
			*resposta = fmt.Sprintf("Conta removida com sucesso e saldo devolvido = %f", a.Saldo)
			s.contas = append(s.contas[:i], s.contas[i+1:]...)
			return nil
		}
	}
	*resposta = "Conta não encontrada."
	return fmt.Errorf("conta com nome %s não encontrada", conta.Nome)
}

func (s *Servidor) Deposito(conta Conta, resposta *string) error {

	for _, a := range s.contas {
		if a.Nome == conta.Nome {
			a.Saldo += conta.Saldo
			*resposta = fmt.Sprintf("Deposito feito, novo de %s saldo = %f", conta.Nome, a.Saldo)
			return nil
		}
	}
	*resposta = "Conta não encontrada."
	return fmt.Errorf("conta com nome %s não encontrada", conta.Nome)
}

func (s *Servidor) Saque(conta Conta, resposta *string) error {

	for _, a := range s.contas {
		if a.Nome == conta.Nome {
			a.Saldo -= conta.Saldo
			*resposta = fmt.Sprintf("Saque feito, novo de %s saldo = %f", conta.Nome, a.Saldo)
			return nil
		}
	}
	*resposta = "Conta não encontrada."
	return fmt.Errorf("conta com nome %s não encontrada", conta.Nome)
}

func main() {
	porta := 8973
	servidor := new(Servidor)
	servidor.inicializar()
	// var resposta string
	// servidor.AbrirConta(struct{}{}, &resposta)

	rpc.Register(servidor)
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", porta))
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
		return
	}

	fmt.Println("Servidor aguardando conexões na porta", porta)
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Erro ao aceitar conexão:", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}
