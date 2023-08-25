package main

import "fmt"

func main() {
  minhaIdade := 19

  maiorDeIdade := minhaIdade >= 18

  if(maiorDeIdade) {
    fmt.Println("Sou maior de idade")
  } else {
    fmt.Println("Vish, é menor de idade")
  }
}

