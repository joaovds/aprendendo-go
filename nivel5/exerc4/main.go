package main

import "fmt"

func main() {
  anonimo := struct{
    palavrasMap map[int]string
    palavras []string
  }{
    palavrasMap: map[int]string{
      1: "Oie",
      2: "tudo bem?",
    },
    palavras: []string{
      "Maça",
      "Maracujá",
    },
  }

  fmt.Println(anonimo)
}

