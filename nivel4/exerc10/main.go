package main

import "fmt"

func main() {
  pessoas := map[string][]string{
    "pessoa 1": {
      "Programar", "Ler", "Estudar Golang",
    },
    "pessoa 2": {
      "Jogar basquete", "Ir a igreja", "Comer",
    },
  }

  delete(pessoas, "pessoa 2")

  for index, pessoa := range pessoas {
    fmt.Printf("Indíce da pessoa %v \n", index)
    fmt.Print("Seus hobbies: ")

    for _, hobbie := range pessoa {
      fmt.Print(hobbie, ", ")
    }

    fmt.Println("\n----------")
  }
}

