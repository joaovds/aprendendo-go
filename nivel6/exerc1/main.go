package main

import "fmt"

type pessoa struct {
  nome string;
  sobrenome string;
  saboresFavoritos []string;
}

func main() {
  apenasInt := retornaInt()
  oInt, aString := retornaIntEString()

  fmt.Println("Apenas int:", apenasInt)
  fmt.Println("Int e string:", oInt, "e", aString)
}

func retornaInt() int {
  return 10;
}

func retornaIntEString() (int, string) {
  return 10, "Dez"
}

