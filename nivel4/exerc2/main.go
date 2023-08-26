package main

import "fmt"

func main() {
  sliceCom10Elementos := []int{ 2, 4, 6, 8, 10, 12, 14, 16, 18, 20 }

  for index, value := range sliceCom10Elementos {
    fmt.Printf("Valor: %v no índice %v\n", value, index)
  }

  fmt.Printf("O tipo da slice é: %T", sliceCom10Elementos)
}

