package main

import "fmt"

var nomeEscola = "Escola Técnica SENAI" 

func main() {
    nome := "Lucas"       
    idade := 16            
    mensagem := boasVindas(nome)
    fmt.Println(mensagem)

    status := verificaMaioridade(idade)
    fmt.Println(status)
}