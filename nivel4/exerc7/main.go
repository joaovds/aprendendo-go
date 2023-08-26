package main

import "fmt"

func main() {
  pessoas := [][]string{
    []string{
      "João",
      "Victor",
      "Programar",
    },
    []string{
      "João 2",
      "Vitor 2",
      "Ler",
    },
    []string{
      "João 3",
      "Vitor 3",
      "Jogar basquete",
    },
  }

  for _, pessoa :=  range pessoas {
    fmt.Println(pessoa)
  }
}

