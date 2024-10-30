package main

import (
    "bufio"
    "fmt"
    "net"
    "net/rpc"
    "os"
    "strconv"
)

// Estrutura para representar um aluno
type Conta struct {
    Nome string
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
		{"Barbara",   855.5},
		{"Joao",      656.5},
		{"Maria",     946.0},
		{"Paulo",    10546.0},
		{"Pedro",     7465.0},
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

// func (s *Servidor) AbrirConta() error {
//     reader := bufio.NewReader(os.Stdin)
// 	fmt.Print("Digite Nome: ")
// 	input, _ := reader.ReadString('\n')
//     Conta conta
//     novo.Nome := input
//     fmt.Print("Digite uma saldo inicial: ")
// 	input, _ := reader.ReadString('\n')
//     novo.Saldo := input
//     s = append(s,novo)
// }

func (s *Servidor) AbrirConta() error {
    reader := bufio.NewReader(os.Stdin)
    // Leitura do nome
    fmt.Print("Digite o nome: ")
    nome, _ := reader.ReadString('\n')
    nome = nome[:len(nome)-1] // Remove o caractere de nova linha
    // Leitura do saldo
    fmt.Print("Digite um saldo inicial: ")
    input, _ := reader.ReadString('\n')
    input = input[:len(input)-1] // Remove o caractere de nova linha
    saldo, err := strconv.ParseFloat(input, 64)
    if err != nil {
        return fmt.Errorf("saldo inválido: %v", err)
    }
    // Criação da nova conta
    conta := Conta{
        Nome:  nome,
        Saldo: saldo,
    }
    // Adiciona a nova conta à lista de contas
    s.contas = append(s.contas, conta)
    fmt.Println("Conta criada com sucesso!")
    return nil
}


func main() {
    porta := 8973
    servidor := new(Servidor)
    servidor.inicializar()

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
