package main

import "fmt"

func main() {
  arrayCom5Elementos := [5]int{ 2, 4, 6, 8, 10 }

  for index, value := range arrayCom5Elementos {
    fmt.Printf("Valor: %v no índice %v\n", value, index)
  }
}

