/*package main

import (
	"fmt"
	"net/rpc"
	"os"
	"sync"
)

var contas_novas = []string{"Bruno", "Sofia", "Izis"}
var contas_antigas = []string{"Maria", "Pedro", "Joao"}

type Conta struct {
	Nome  string
	Saldo float64
}

// Estrutura que mantém o cliente RPC
type RPCClient struct {
	client *rpc.Client
	mutex  sync.Mutex
}

// Função para criar uma nova instância de RPCClient e abrir a conexão
func NewRPCClient(maquina string, porta int) (*RPCClient, error) {
	client, err := rpc.Dial("tcp", fmt.Sprintf("%s:%d", maquina, porta))
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar ao servidor: %v", err)
	}
	return &RPCClient{client: client}, nil
}

// Método para fechar a conexão
func (rpcClient *RPCClient) Close() {
	rpcClient.client.Close()
}

// Método para abrir uma conta
func (rpcClient *RPCClient) ABRIR(nome string) {
	var resposta string
	err := rpcClient.client.Call("Servidor.AbrirConta", nome, &resposta)
	if err != nil {
		fmt.Println("Erro ao abrir conta:", err)
	} else {
		fmt.Println("Resposta do servidor:", resposta)
	}
}

// Método para fechar uma conta
func (rpcClient *RPCClient) FECHAR(nome string) {
	var resposta string
	err := rpcClient.client.Call("Servidor.FecharConta", nome, &resposta)
	if err != nil {
		fmt.Println("Erro ao fechar conta:", err)
	} else {
		fmt.Println("Resposta do servidor:", resposta)
	}
}

// Método para realizar um depósito
func (rpcClient *RPCClient) DEPOSITO(conta Conta) {
	var resposta string
	err := rpcClient.client.Call("Servidor.Deposito", conta, &resposta)
	if err != nil {
		fmt.Println("Erro no depósito:", err)
	} else {
		fmt.Println("Resposta do servidor:", resposta)
	}
}

// Método para realizar um saque
func (rpcClient *RPCClient) SAQUE(conta Conta) {
	var resposta string
	err := rpcClient.client.Call("Servidor.Saque", conta, &resposta)
	if err != nil {
		fmt.Println("Erro no saque:", err)
	} else {
		fmt.Println("Resposta do servidor:", resposta)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Uso:", os.Args[0], "<maquina>")
		return
	}
	porta := 8973
	maquina := os.Args[1]

	// Criação do RPCClient
	rpcClient, err := NewRPCClient(maquina, porta)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rpcClient.Close()

	// Abrindo contas
	for _, nome := range contas_novas {
		rpcClient.ABRIR(nome)
	}

	// Fazendo depósitos e saques
	var wg sync.WaitGroup
	for _, nome := range contas_novas {
		wg.Add(1)
		go func(nome string) {
			defer wg.Done()
			rpcClient.DEPOSITO(Conta{Nome: nome, Saldo: 500.0})
		}(nome)
	}

	for _, nome := range contas_antigas {
		wg.Add(1)
		go func(nome string) {
			defer wg.Done()
			rpcClient.SAQUE(Conta{Nome: nome, Saldo: 200.0})
		}(nome)
	}

	wg.Wait() // Aguarda todas as goroutines terminarem

	// Fechando contas
	for _, nome := range contas_antigas {
		rpcClient.FECHAR(nome)
	}
}*/