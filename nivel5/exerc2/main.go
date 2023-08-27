package main

import "fmt"

type pessoa struct {
  nome string;
  sobrenome string;
  saboresFavoritos []string;
}

func main() {
  pessoas := make(map[string]pessoa);

  pessoas["João"] = pessoa{
    nome: "João",
    sobrenome: "Victor",
    saboresFavoritos: []string{
      "Sensação",
      "Floresta Negra",
      "Abacaxi com vinho",
    },
  }

  pessoas["Nicoly"] = pessoa{
    nome: "Nicoly",
    sobrenome: "Araújo",
    saboresFavoritos: []string{
      "Ninho",
      "Maracujá",
      "Todos hahah",
    },
  }

  for _, pessoa := range pessoas {
    fmt.Println(pessoa.nome, pessoa.sobrenome)
    for _, sabor := range pessoa.saboresFavoritos {
      fmt.Print(sabor, " - ")
    }
    fmt.Println(" ")
  }
}

