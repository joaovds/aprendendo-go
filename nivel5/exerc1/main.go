package main

import "fmt"

type pessoa struct {
  nome string;
  sobrenome string;
  saboresFavoritos []string;
}

func main() {
  pessoa1 := pessoa{
    nome: "João",
    sobrenome: "Victor",
    saboresFavoritos: []string{
      "Sensação",
      "Floresta Negra",
      "Abacaxi com vinho",
    },
  }

  pessoa2 := pessoa{
    nome: "Nicoly",
    sobrenome: "Araújo",
    saboresFavoritos: []string{
      "Ninho",
      "Maracujá",
      "Todos hahah",
    },
  }

  fmt.Println(pessoa1.nome, pessoa1.sobrenome)
  for _, sabor := range pessoa1.saboresFavoritos {
    fmt.Print(sabor, " - ")
  }

  fmt.Println(" ")

  fmt.Println(pessoa2.nome, pessoa2.sobrenome)
  for _, sabor := range pessoa2.saboresFavoritos {
    fmt.Print(sabor, " - ")
  }
}

