package main

import (
	"fmt"
	"net/rpc"
	"os"
	"sync"
)

var contas_novas = []string{"Bruno", "Sofia", "Izis"}
var contas_antigas = []string{"Maria", "Pedro", "Joao"}

var wg sync.WaitGroup

type Conta struct {
	Nome  string
	Saldo float64
}

func SALDO(nome string, porta int, maquina string, cliente *rpc.Client) {
	var resposta string
	var err = cliente.Call("Servidor.ObtemSaldo", nome, &resposta)
	if err != nil {
		fmt.Println("Erro ao ver saldo conta:", err)
	} else {
		fmt.Println("Resposta do servidor", resposta)
	}

}

func ABRIR(nome string, porta int, maquina string, cliente *rpc.Client) {
	var resposta string
	var err = cliente.Call("Servidor.AbrirConta", nome, &resposta)
	if err != nil {
		fmt.Println("Erro ao abrir conta:", err)
	} else {
		fmt.Println("Resposta do servidor:", resposta)
	}

}
func FECHAR(nome string, porta int, maquina string, client *rpc.Client) {
	var resposta string
	var err = client.Call("Servidor.FecharConta", nome, &resposta)
	if err != nil {
		fmt.Println("Erro ao fechar conta:", err)
	} else {
		fmt.Println("Resposta do servidor:", resposta)
	}

}

func DEPOSITO(nome string, porta int, maquina string, client *rpc.Client) {
	var resposta string
	var err = client.Call("Servidor.Deposito", Conta{Nome: nome, Saldo: 200}, &resposta)
	if err != nil {
		fmt.Println("Erro no deposito conta:", err)
	} else {
		fmt.Println("Resposta do servidor:", resposta)
	}
	wg.Done()

}
func SAQUE(nome string, porta int, maquina string, client *rpc.Client) {
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
	porta := 8973
	maquina := os.Args[1]
	client, err := rpc.Dial("tcp", fmt.Sprintf("%s:%d", maquina, porta))
	if err != nil {
		fmt.Println("Erro ao conectar ao servidor:", err)
		return
	}
	defer client.Close()

	// for i := 0; i < 3; i++ {
	// 	ABRIR(contas_novas[i], porta, maquina, client)
	// }
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go DEPOSITO(contas_novas[i], porta, maquina, client)
	}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go SAQUE(contas_antigas[i], porta, maquina, client)
	}
	wg.Wait()
	for i := 0; i < 3; i++ {
		SALDO(contas_antigas[i], porta, maquina, client)
	}

}