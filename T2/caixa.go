package main

import (
	"fmt"
	"net/rpc"
	"os"
	"sync"
)

var contas_novas = []string{"Bruno", "Sofia", "Izis", "Enzo", "Carlos", "Gabriel"}
var contas_antigas = []string{"Maria", "Pedro", "Joao", "Alexandre", "Barbara", "Paulo"}

// Estrurura wait para Goroutines
var wg sync.WaitGroup

type Conta struct {
	Nome  string
	Saldo float64
}

// Metodo de verificação do saldo disponivel para um nome
func SALDO(nome string, cliente *rpc.Client) {
	var resposta string
	var err = cliente.Call("Servidor.ObtemSaldo", nome, &resposta)
	if err != nil {
		fmt.Println("Erro ao ver saldo conta:", err)
	} else {
		fmt.Println("Resposta do servidor:", resposta)
	}
	wg.Done()
}

// Metodo de Abrir conta
func ABRIR(nome string, cliente *rpc.Client) {
	var resposta string
	var err = cliente.Call("Servidor.AbrirConta", nome, &resposta)
	if err != nil {
		fmt.Println("Erro ao abrir conta:", err)
	} else {
		fmt.Println("Resposta do servidor:", resposta)
	}
	wg.Done()

}

// Metodo de Fechar  conta
func FECHAR(nome string, client *rpc.Client) {
	var resposta string
	var err = client.Call("Servidor.FecharConta", nome, &resposta)
	if err != nil {
		fmt.Println("Erro ao fechar conta:", err)
	} else {
		fmt.Println("Resposta do servidor:", resposta)
	}
	wg.Done()

}

// Metodo de Depositar em conta
func DEPOSITO(nome string, client *rpc.Client) {
	var resposta string
	var err = client.Call("Servidor.Deposito", Conta{Nome: nome, Saldo: 200}, &resposta)
	if err != nil {
		fmt.Println("Erro no deposito conta:", err)
	} else {
		fmt.Println("Resposta do servidor:", resposta)
	}
	wg.Done()

}

// Metodo de saque em conta
func SAQUE(nome string, client *rpc.Client) {
	var resposta string
	var err = client.Call("Servidor.Saque", Conta{Nome: nome, Saldo: 100}, &resposta)
	if err != nil {
		fmt.Println("Erro no saque conta:", err)
	} else {
		fmt.Println("Resposta do servidor:", resposta)
	}
	wg.Done()

}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Uso:", os.Args[0], "<maquina>")
		return
	}
	porta := 1234
	maquina := os.Args[1]
	client, err := rpc.Dial("tcp", fmt.Sprintf("%s:%d", maquina, porta))
	if err != nil {
		fmt.Println("Erro ao conectar ao servidor:", err)
		return
	}
	defer client.Close()

	for i := 0; i < len(contas_antigas); i++ {
		wg.Add(1)
		go DEPOSITO(contas_novas[i], client)
		wg.Add(1)
		go DEPOSITO(contas_antigas[i], client)
	}
	for i := 0; i < len(contas_antigas); i++ {
		wg.Add(1)
		go SAQUE(contas_antigas[i], client)
		wg.Add(1)
		go SAQUE(contas_novas[i], client)
	}
	wg.Wait()
	for i := 0; i < len(contas_antigas); i++ {
		wg.Add(1)
		go SALDO(contas_antigas[i], client)
		wg.Add(1)
		go SALDO(contas_novas[i], client)
	}
	wg.Wait()

}
