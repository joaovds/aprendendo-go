package main

import "fmt"

type Pessoa struct {
  nome string;
}

func main() {
  pessoa := Pessoa{
    nome: "João",
  }

  fmt.Println("Pessoa antes: ", pessoa)

  mudeMe(&pessoa)

  fmt.Println("Pessoa depois: ", pessoa)
}

func mudeMe(pessoa *Pessoa) {
  (*pessoa).nome = "João 2, pq mudou aqui com ponteiro"
}

