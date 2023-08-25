package main

import "fmt"

func main() {
  anoNascimento := 2004
  anoAtual := 2023

  for anoNascimento <= anoAtual {
    fmt.Println(anoNascimento)

    anoNascimento++
  }
}

