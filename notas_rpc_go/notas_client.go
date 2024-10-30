// package main

// import (
//     "fmt"
//     "net/rpc"
//     "os"
// )

// func main() {

//     if len(os.Args) != 3 {
//         fmt.Println("Uso:", os.Args[0], " <maquina> <nome_do_aluno>")
//         return
//     }

//     porta := 8973
//     maquina := os.Args[1]
//     nome := os.Args[2]

//     client, err := rpc.Dial("tcp", fmt.Sprintf("%s:%d", maquina, porta))
//     if err != nil {
//         fmt.Println("Erro ao conectar ao servidor:", err)
//         return
//     }
//     i int := 2
    
//     var nota float64
//     err = client.Call("Servidor.ObtemSaldo", nome, &nota)
//     while (i < 5){
//         if err != nil {
//             fmt.Println("Erro ao obter nota:", err)
//         } else {
//             fmt.Printf("Nome: %s\n", nome)
//             fmt.Printf("Saldo: %.2f\n", nota)
//         }
//         i := i + 1
//     }
// }


package main

import (
    "fmt"
    "net/rpc"
    "os"
    "bufio"
)

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
        fmt.Println("3 - Sair")
        fmt.Print("Escolha uma opção: ")
        opcao, _ := reader.ReadString('\n')
        opcao = opcao[:len(opcao)-1]

        switch opcao {
        case "1":
            // Chamada para AbrirConta
            var resposta string
            err := client.Call("Servidor.AbrirConta", struct{}{}, &resposta)
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

            var saldo float64
            err := client.Call("Servidor.ObtemSaldo", nome, &saldo)
            if err != nil {
                fmt.Println("Erro ao obter saldo:", err)
            } else {
                fmt.Printf("Saldo de %s: %.2f\n", nome, saldo)
            }

        case "3":
            fmt.Println("Saindo...")
            return

        default:
            fmt.Println("Opção inválida. Tente novamente.")
        }
    }
}
