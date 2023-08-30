package main

import "fmt"

type pessoa struct {
  nome string;
  sobrenome string;
  saboresFavoritos []string;
}

func main() {
  soma := somaInts(1, 2, 3, 4, 5)
  somaSlice := somaIntsSlice([]int{1, 2, 3, 4, 5})

  fmt.Println("A soma dos ints:", soma)
  fmt.Println("A soma dos ints do slice:", somaSlice)
}

func somaInts(inteiros ...int) int {
  total := 0

  for _, inteiro := range inteiros {
    total += inteiro
  }

  return total
}

func somaIntsSlice(inteiros []int) int {
  total := 0

  for _, inteiro := range inteiros {
    total += inteiro
  }

  return total
}

