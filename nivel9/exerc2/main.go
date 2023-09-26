package main

import (
  "fmt"
)

type Pessoa struct {
  nome string
}

type Humanos interface {
  falar()
}

func (p *Pessoa) falar() {
  fmt.Println("Oi, meu nome é", p.nome)
}

func dizerAlgumaCoisa(h Humanos) {
  h.falar()
}

func main() {
  pessoa := Pessoa{nome: "João"}

  (&pessoa).falar()
  dizerAlgumaCoisa(&pessoa)
}

