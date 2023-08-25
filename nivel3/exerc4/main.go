package main

import "fmt"

func main() {
  anoNascimento := 2004
  anoAtual := 2023

  for {
    fmt.Println(anoNascimento)

    anoNascimento++

    if(anoNascimento == anoAtual) { break }
  }
}

