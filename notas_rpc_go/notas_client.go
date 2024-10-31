package main

import (
	"bufio"
	"fmt"
	"net/rpc"
	"os"
	"strconv"
)

type Conta struct {
	Nome  string
	Saldo float64
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

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1 - Abrir Conta")
		fmt.Println("2 - Obter Saldo")
		fmt.Println("3 - Fechar Conta")
		fmt.Println("4 - Deposito")
		fmt.Println("5 - Saque ")
		fmt.Println("6 - Sair\n")
		fmt.Print("Escolha uma opção: ")
		opcao, _ := reader.ReadString('\n')
		opcao = opcao[:len(opcao)-1]

		switch opcao {
		case "1":
			// Chamada para AbrirConta
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
				fmt.Println("Saldo inválido:", err)
				continue
			}
			conta := Conta{
				Nome:  nome,
				Saldo: saldo,
			}

			var resposta string
			err = client.Call("Servidor.AbrirConta", conta, &resposta)
			if err != nil {
				fmt.Println("Erro ao abrir conta:", err)
			} else {
				fmt.Println("Resposta do servidor:", resposta)
			}

		case "2":
			// Solicita o nome do cliente e obtém o saldo
			fmt.Print("Digite o nome do cliente: ")
			nome, _ := reader.ReadString('\n')
			nome = nome[:len(nome)-1]

			var saldo_verify float64
			err := client.Call("Servidor.ObtemSaldo", nome, &saldo_verify)
			if err != nil {
				fmt.Println("Erro ao obter saldo:", err)
			} else {
				fmt.Printf("Saldo de %s: %.2f\n", nome, saldo_verify)
			}
		case "3":
			fmt.Print("Digite o nome: ")
			nome, _ := reader.ReadString('\n')
			nome = nome[:len(nome)-1] // Remove o caractere de nova linha
			var saldo float64 = 0
			conta := Conta{
				Nome:  nome,
				Saldo: saldo,
			}
			var resposta string
			err = client.Call("Servidor.FecharConta", conta, &resposta)
			if err != nil {
				fmt.Println("Erro ao Fechar conta:", err)
			} else {
				fmt.Println("Resposta do servidor:", resposta)
			}

		case "4":
			fmt.Print("Digite o nome: ")
			nome, _ := reader.ReadString('\n')
			nome = nome[:len(nome)-1] // Remove o caractere de nova linha

			// Leitura do saldo
			fmt.Print("Digite o deposito: ")
			input, _ := reader.ReadString('\n')
			input = input[:len(input)-1] // Remove o caractere de nova linha
			saldo, err := strconv.ParseFloat(input, 64)
			if err != nil || saldo < 0 {
				fmt.Println("Saldo inválido:", err)
				continue
			}
			conta := Conta{
				Nome:  nome,
				Saldo: saldo,
			}

			var resposta string
			err = client.Call("Servidor.Deposito", conta, &resposta)
			if err != nil {
				fmt.Println("Erro ao abrir conta:", err)
			} else {
				fmt.Println("Resposta do servidor:", resposta)
			}

		case "5":
			fmt.Print("Digite o nome: ")
			nome, _ := reader.ReadString('\n')
			nome = nome[:len(nome)-1] // Remove o caractere de nova linha

			// Leitura do saldo
			fmt.Print("Digite o saque: ")
			input, _ := reader.ReadString('\n')
			input = input[:len(input)-1] // Remove o caractere de nova linha
			saldo, err := strconv.ParseFloat(input, 64)
			if err != nil || saldo < 0 {
				fmt.Println("Saldo inválido:", err)
				continue
			}
			conta := Conta{
				Nome:  nome,
				Saldo: saldo,
			}

			var resposta string
			err = client.Call("Servidor.Saque", conta, &resposta)
			if err != nil {
				fmt.Println("Erro ao abrir conta:", err)
			} else {
				fmt.Println("Resposta do servidor:", resposta)
			}

		case "6":
			fmt.Println("Saindo...")
			return

		default:
			fmt.Println("Opção inválida. Tente novamente.")
		}
	}
}
