package main

import "fmt"

func main() {
  minhaIdade := 19

  maiorDeIdade := minhaIdade >= 18 && minhaIdade < 60
  idoso := minhaIdade >= 60

  if maiorDeIdade {
    fmt.Println("Sou maior de idade")
  } else if idoso {
    fmt.Println("Sou um idoso")
  }else {
    fmt.Println("Vish, é menor de idade")
  }
}

