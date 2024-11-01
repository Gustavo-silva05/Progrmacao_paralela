package main

import (
	"fmt"
	"net/rpc"
	"os"
	"sync"
)

var contas_novas = []string{"Bruno", "Sofia", "Izis"}
var contas_antigas = []string{"Maria", "Pedro", "Joao"}

var mutex sync.Mutex
var wg sync.WaitGroup

type Conta struct {
	Nome  string
	Saldo float64
}

func ABRIR(nome string) {
	porta := 8973
	maquina := os.Args[1]
	client, err := rpc.Dial("tcp", fmt.Sprintf("%s:%d", maquina, porta))
	if err != nil {
		fmt.Println("Erro ao conectar ao servidor:", err)
		return
	}
	defer client.Close()
	var resposta string
	err = client.Call("Servidor.AbrirConta", nome, &resposta)
	if err != nil {
		fmt.Println("Erro ao abrir conta:", err)
	} else {
		fmt.Println("Resposta do servidor:", resposta)
	}

}
func FECHAR(nome string) {
	porta := 8973
	maquina := os.Args[1]
	client, err := rpc.Dial("tcp", fmt.Sprintf("%s:%d", maquina, porta))
	if err != nil {
		fmt.Println("Erro ao conectar ao servidor:", err)
		return
	}
	defer client.Close()
	var resposta string
	err = client.Call("Servidor.FecharConta", nome, &resposta)
	if err != nil {
		fmt.Println("Erro ao fechar conta:", err)
	} else {
		fmt.Println("Resposta do servidor:", resposta)
	}

}

func DEPOSITO(nome string, id int) {
	porta := 8973
	maquina := os.Args[1]
	conta := Conta{Nome: nome, Saldo: 500.0}
	client, err := rpc.Dial("tcp", fmt.Sprintf("%s:%d", maquina, porta))
	if err != nil {
		fmt.Println("Erro ao conectar ao servidor:", err)
		return
	}
	defer client.Close()
	var resposta string
	err = client.Call("Servidor.Deposito", conta, &resposta)
	if err != nil {
		fmt.Println("Erro no deposito conta:", err)
	} else {
		fmt.Println("Resposta do servidor:", resposta)
	}

}
func SAQUE(nome string, id int) {
	porta := 8973
	maquina := os.Args[1]
	conta := Conta{Nome: nome, Saldo: 200}
	client, err := rpc.Dial("tcp", fmt.Sprintf("%s:%d", maquina, porta))
	if err != nil {
		fmt.Println("Erro ao conectar ao servidor:", err)
		return
	}
	defer client.Close()
	var resposta string
	err = client.Call("Servidor.Saque", conta, &resposta)
	if err != nil {
		fmt.Println("Erro no saque conta:", err)
	} else {
		fmt.Println("Resposta do servidor:", resposta)
	}

}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Uso:", os.Args[0], "<maquina>")
		return
	}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go ABRIR(contas_novas[i])
	}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go DEPOSITO(contas_novas[i], i+1)
	}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go SAQUE(contas_antigas[i], i+1)
	}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go FECHAR(contas_antigas[i])
	}
	wg.Done()
	wg.Wait()

}
