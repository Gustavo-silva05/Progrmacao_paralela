package main

import (
	"fmt"
	"net"
	"net/rpc"
	"sync"
)

var mutex sync.Mutex

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

func (s *Servidor) AbrirConta(nome string, resposta *string) error {
	conta := Conta{
		Nome:  nome,
		Saldo: 0.0,
	}
	s.contas = append(s.contas, conta)
	fmt.Println("Conta criada para ", nome)
	*resposta = fmt.Sprintf("Conta de %s criada com sucesso!", nome)
	return nil
}

func (s *Servidor) FecharConta(nome string, resposta *string) error {

	for i, a := range s.contas {
		if a.Nome == nome {
			mutex.Lock()
			// Remove a conta da lista
			fmt.Println("Conta excluida de ", nome)
			*resposta = fmt.Sprintf("Conta removida com sucesso e saldo devolvido = %g", a.Saldo)
			s.contas = append(s.contas[:i], s.contas[i+1:]...)
			mutex.Unlock()
			return nil

		}
	}
	*resposta = "Conta não encontrada."
	return fmt.Errorf("conta com nome %s não encontrada", nome)
}

func (s *Servidor) Deposito(conta Conta, resposta *string) error {

	for i, a := range s.contas {
		if a.Nome == conta.Nome {
			mutex.Lock()
			s.contas[i].Saldo += conta.Saldo
			fmt.Println("Desposito Realizado com sucesso ")
			*resposta = fmt.Sprintf("Deposito de %g feito, novo saldo de %s = %g", conta.Saldo, conta.Nome, s.contas[i].Saldo)
			mutex.Unlock()
			return nil
		}
	}
	*resposta = "Conta não encontrada."
	return fmt.Errorf("conta com nome %s não encontrada", conta.Nome)
}

func (s *Servidor) Saque(conta Conta, resposta *string) error {

	for i, a := range s.contas {
		if a.Nome == conta.Nome {
			mutex.Lock()
			s.contas[i].Saldo -= conta.Saldo
			fmt.Println("Saque Realizado com sucesso ")
			*resposta = fmt.Sprintf("Saque de %g feito, novo saldo de %s = %g", conta.Saldo, conta.Nome, s.contas[i].Saldo)
			mutex.Unlock()
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
