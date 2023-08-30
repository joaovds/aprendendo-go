package main

import "fmt"

type pessoa struct {
  nome string;
  sobrenome string;
  idade int;
}

func main() {
  umaPessoa := pessoa{
    nome: "João",
    sobrenome: "Silva",
    idade: 19,
  }

  umaPessoa.mostrarInfosPessoa()
}

func (p pessoa) mostrarInfosPessoa() {
  fmt.Println("Nome:", p.nome, p.sobrenome, "\t idade", p.idade)
}

